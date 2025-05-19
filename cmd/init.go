package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Get path to an executable in $PATH
func getExecutablePath(executable string) string {
	cmd := exec.Command("which", executable)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error finding %s in PATH: %s\n", executable, err)
		return ""
	}
	return strings.TrimSpace(string(out))
}

// Login command (optional, if not already handled)
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Infra",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load(".env")
		infraUser := os.Getenv("INFRA_USER")
		infraPassword := os.Getenv("INFRA_PASSWORD")
		infraPath := getExecutablePath("infra")

		if infraPath == "" || infraUser == "" || infraPassword == "" {
			fmt.Println("Missing infra executable or credentials")
			return
		}

		command := exec.Command(infraPath, "login", "infra-k8s.supervity.ai", "--user", infraUser)
		stdin, err := command.StdinPipe()
		if err != nil {
			fmt.Println("Error creating stdin pipe:", err)
			return
		}

		if err := command.Start(); err != nil {
			fmt.Println("Error starting infra login:", err)
			return
		}

		_, _ = stdin.Write([]byte(infraPassword + "\n"))
		_ = stdin.Close()
		if err := command.Wait(); err != nil {
			fmt.Println("Infra login failed:", err)
			return
		}

		fmt.Println("✅ Logged in to Infra")
	},
}

// Init command: list clusters, select, fetch pods, select app, stream logs
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "List clusters and stream logs for selected app",
	Run: func(cmd *cobra.Command, args []string) {
		infra := getExecutablePath("infra")
		if infra == "" {
			fmt.Println("infra not found")
			return
		}

		// List clusters
		listCmd := exec.Command(infra, "list")
		out, err := listCmd.CombinedOutput()
		if err != nil {
			fmt.Println("infra list failed:", err)
			fmt.Println(string(out))
			return
		}

		lines := strings.Split(string(out), "\n")
		var clusters []string
		for _, line := range lines {
			if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "NAME") {
				clusters = append(clusters, strings.Fields(line)[0])
			}
		}

		clusterPrompt := promptui.Select{
			Label: "Select Cluster",
			Items: clusters,
		}
		_, cluster, err := clusterPrompt.Run()
		if err != nil {
			fmt.Println("Cluster selection failed:", err)
			return
		}

		// Select cluster
		useCmd := exec.Command(infra, "use", cluster)
		useOut, err := useCmd.CombinedOutput()
		if err != nil {
			fmt.Println("infra use failed:", err)
			fmt.Println(string(useOut))
			return
		}

		listPods()
	},
}

func listPods() {
	kubectl := getExecutablePath("kubectl")
	if kubectl == "" {
		fmt.Println("kubectl not found")
		return
	}

	// Get pods with labels
	cmd := exec.Command(kubectl, "get", "pods", "--show-labels")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("kubectl get pods failed:", err)
		fmt.Println(string(out))
		return
	}

	lines := strings.Split(string(out), "\n")
	appLabels := map[string]bool{}
	for _, line := range lines {
		if strings.HasPrefix(line, "NAME") || strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		labelCol := fields[len(fields)-1]
		labels := strings.Split(labelCol, ",")
		for _, lbl := range labels {
			lbl = strings.TrimSpace(lbl)
			if strings.HasPrefix(lbl, "app=") {
				appLabels[lbl] = true
			}
		}
	}

	// Convert map to slice
	var labelList []string
	for k := range appLabels {
		labelList = append(labelList, k)
	}

	labelPrompt := promptui.Select{
		Label: "Select app label",
		Items: labelList,
	}
	_, selectedLabel, err := labelPrompt.Run()
	if err != nil {
		fmt.Println("Label selection failed:", err)
		return
	}

	getLogs(selectedLabel)
}

func getLogs(label string) {
	kubectl := getExecutablePath("kubectl")
	if kubectl == "" {
		fmt.Println("kubectl not found")
		return
	}

	cmd := exec.Command(kubectl, "logs", "-f", "-l", label, "--since=1h", "--prefix")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting logs command:", err)
		return
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error running kubectl logs: %s\n", err)
	}
}
