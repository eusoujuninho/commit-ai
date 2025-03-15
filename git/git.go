// Package git fornece funcionalidades para interação com repositórios Git
package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

// Repository representa um repositório Git
type Repository struct {
	Path string
	repo *git.Repository
}

// OpenRepository abre um repositório Git existente
func OpenRepository(path string) (*Repository, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir repositório: %w", err)
	}

	return &Repository{
		Path: path,
		repo: repo,
	}, nil
}

// GetChangedFiles retorna uma lista dos arquivos modificados
func (r *Repository) GetChangedFiles() ([]string, error) {
	w, err := r.repo.Worktree()
	if err != nil {
		return nil, err
	}

	status, err := w.Status()
	if err != nil {
		return nil, err
	}

	var changedFiles []string
	for file, fileStatus := range status {
		if fileStatus.Staging != git.Unmodified || fileStatus.Worktree != git.Unmodified {
			changedFiles = append(changedFiles, file)
		}
	}

	return changedFiles, nil
}

// GetChanges retorna as diferenças dos arquivos modificados
func (r *Repository) GetChanges() (string, error) {
	// Obtém a lista de arquivos modificados
	files, err := r.GetChangedFiles()
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "Nenhuma alteração detectada", nil
	}

	// Construir resumo das alterações
	summary := fmt.Sprintf("Alterações em %d arquivo(s):\n", len(files))
	for _, file := range files {
		summary += fmt.Sprintf("- %s\n", file)
	}
	summary += "\nDetalhes das alterações:\n"

	// Obter status detalhado das alterações
	cmdStatus := exec.Command("git", "-C", r.Path, "status", "-v")
	statusOutput, err := cmdStatus.Output()
	if err == nil {
		summary += "\nStatus do repositório:\n" + string(statusOutput) + "\n"
	}

	// Obter informações sobre os arquivos alterados (extensão, tipo, etc.)
	fileInfos := make(map[string]string)
	for _, file := range files {
		// Verificar tipo de arquivo
		fileType := "arquivo"
		if strings.HasSuffix(file, ".go") {
			fileType = "código Go"
		} else if strings.HasSuffix(file, ".md") {
			fileType = "documentação (Markdown)"
		} else if strings.HasSuffix(file, ".js") {
			fileType = "código JavaScript"
		} else if strings.HasSuffix(file, ".html") {
			fileType = "arquivo HTML"
		} else if strings.HasSuffix(file, ".css") {
			fileType = "estilo CSS"
		} else if strings.HasSuffix(file, ".json") {
			fileType = "dados JSON"
		} else if strings.HasSuffix(file, ".yml") || strings.HasSuffix(file, ".yaml") {
			fileType = "configuração YAML"
		}
		fileInfos[file] = fileType
	}

	fileContext := "\nContexto dos arquivos:\n"
	for file, info := range fileInfos {
		fileContext += fmt.Sprintf("- %s: %s\n", file, info)
	}
	summary += fileContext

	// Executar git diff para obter as diferenças reais
	hasStagedChanges, _ := hasStaged(r.Path)

	// Obter diff mais detalhado com contexto adicional
	var diffCmd *exec.Cmd
	if hasStagedChanges {
		// Obter diff das alterações staged com mais contexto
		diffCmd = exec.Command("git", "-C", r.Path, "diff", "--cached", "--no-color", "-U10")
	} else {
		// Obter diff das alterações não staged com mais contexto
		diffCmd = exec.Command("git", "-C", r.Path, "diff", "--no-color", "-U10")
	}

	output, err := diffCmd.Output()
	if err != nil {
		// Se falhar, tentar com menos opções
		diffCmd = exec.Command("git", "diff", "--no-color")
		diffCmd.Dir = r.Path
		output, err = diffCmd.Output()
		if err != nil {
			// Se ainda falhar, retornar apenas o resumo
			return summary + "\nNão foi possível obter o diff detalhado.", nil
		}
	}

	// Obter estatísticas das alterações
	statsCmd := exec.Command("git", "-C", r.Path, "diff", "--stat")
	statsOutput, err := statsCmd.Output()
	if err == nil {
		summary += "\nEstatísticas das alterações:\n" + string(statsOutput) + "\n"
	}

	// Aumentar o limite do diff para capturar mais contexto
	diff := string(output)
	maxDiffSize := 8000 // Aumentado para obter mais contexto
	if len(diff) > maxDiffSize {
		diff = diff[:maxDiffSize] + "...\n[Diff truncado devido ao tamanho]"
	}

	return summary + "\nDiff completo:\n" + diff, nil
}

// hasStaged verifica se há alterações staged no repositório
func hasStaged(repoPath string) (bool, error) {
	cmd := exec.Command("git", "-C", repoPath, "diff", "--cached", "--quiet")
	err := cmd.Run()
	// Se o comando retornar código diferente de zero, há alterações staged
	return err != nil, nil
}

// Commit realiza um commit com a mensagem especificada
func (r *Repository) Commit(message string) error {
	w, err := r.repo.Worktree()
	if err != nil {
		return err
	}

	// Verificar se há alterações staged
	hasStagedChanges, _ := hasStaged(r.Path)
	if !hasStagedChanges {
		// Se não houver alterações staged, adicionar todos os arquivos modificados
		changedFiles, err := r.GetChangedFiles()
		if err != nil {
			return err
		}

		// Adicionar cada arquivo modificado ao stage
		for _, file := range changedFiles {
			_, err = w.Add(file)
			if err != nil {
				return fmt.Errorf("erro ao adicionar arquivo %s: %w", file, err)
			}
		}
	}

	// Realizar o commit
	_, err = w.Commit(message, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}

// AddFilesToStage adiciona uma lista de arquivos para a área de stage
func (r *Repository) AddFilesToStage(files []string) error {
	w, err := r.repo.Worktree()
	if err != nil {
		return err
	}

	// Adicionar cada arquivo ao stage
	for _, file := range files {
		_, err = w.Add(file)
		if err != nil {
			return fmt.Errorf("erro ao adicionar arquivo %s: %w", file, err)
		}
	}

	return nil
}
