package usecase

import (
	"github.com/VadimRight/GraficalTreeWork/entities"
)

type BuildTreeUseCase struct {
	Tree *entities.BSTFuzzy
}

func NewBuildTreeUseCase() *BuildTreeUseCase {
	return &BuildTreeUseCase{Tree: &entities.BSTFuzzy{}}
}

func (uc *BuildTreeUseCase) InsertNode(value int, membership float64) {
	uc.Tree.InsertFuzzy(uc.Tree.Root, value, membership)
}

func (uc *BuildTreeUseCase) SearchNode(value int, membership float64) bool {
	return uc.Tree.SearchFuzzy(uc.Tree.Root, value, membership)
}
