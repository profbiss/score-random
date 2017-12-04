package score_random

import (
	"math/rand"
	"time"
)

type Item struct{
	Id int
	Score float32
	weight float32
	suggests float32
}

func NewItem(id int, score float32) Item {
	return Item{id,score, 0, 0}
}

type Suggest struct {
	Items []Item
	sumScores float32
	sumWeights float32
	sumSuggests float32
	random *rand.Rand
}

func NewSuggest() Suggest {
	suggest := Suggest{}
	suggest.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	return suggest
}

func (self *Suggest) calculateWeights () {
	self.sumScores = 0
	self.sumWeights = 0
	for _, item := range self.Items {
		self.sumScores += item.Score
	}
	for i, item := range self.Items {
		self.Items[i].weight = item.Score / self.sumScores
		self.sumWeights += self.Items[i].weight
	}
}

func (self *Suggest) GetItem() *Item {
	self.calculateWeights()
	randomResult := self.random.Float32()

	var cursor float32
	for i, item := range self.Items {
		cursor += item.weight / self.sumWeights
		if cursor >= randomResult {
			self.Items[i].suggests++
			self.sumSuggests++
			return &item
		}
	}
	return nil
}