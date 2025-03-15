// Package git fornece funcionalidades para interação com repositórios Git
package git

import (
	"fmt"
	"os/exec"

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

	// Executar git diff para obter as diferenças reais
	cmd := exec.Command("git", "-C", r.Path, "diff", "--cached", "--no-color")
	// Se não houver mudanças staged, mostrar as mudanças não staged
	hasStagedChanges, _ := hasStaged(r.Path)
	if !hasStagedChanges {
		cmd = exec.Command("git", "-C", r.Path, "diff", "--no-color")
	}

	output, err := cmd.Output()
	if err != nil {
		// Se falhar, tentar obter diff sem o parâmetro -C
		cmd = exec.Command("git", "diff", "--no-color")
		cmd.Dir = r.Path
		output, err = cmd.Output()
		if err != nil {
			// Se ainda falhar, retornar apenas o resumo
			return summary + "\nNão foi possível obter o diff detalhado.", nil
		}
	}

	// Se o diff for muito grande, limitar para não sobrecarregar o modelo
	diff := string(output)
	maxDiffSize := 4000 // Limitar para não exceder o contexto do modelo
	if len(diff) > maxDiffSize {
		diff = diff[:maxDiffSize] + "...\n[Diff truncado devido ao tamanho]"
	}

	return summary + diff, nil
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
