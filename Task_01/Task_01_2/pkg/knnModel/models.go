package knnModel

import (
	"errors"
	"math"
	"sort"
)

type Entity struct {
	ClassName       string
	AttributeValues []float64
}

type TrainingSet []Entity

type Metric func(Entity, Entity) float64

type KnnModel struct {
	_metric         Metric
	_set            TrainingSet
	_attributeCount int
	K               int
}

func NewKnnModel(m Metric, n int, k int) *KnnModel {
	return &KnnModel{
		_metric:         m,
		_attributeCount: n,
		K:               k,
		_set:            make([]Entity, 0),
	}
}

func (model *KnnModel) AddEntities(data ...Entity) {
	for _, entry := range data {
		model._set = append(model._set, entry)
	}
}

func (model *KnnModel) Classify(attrValues ...float64) (string, error) {
	if model == nil || len(attrValues) < model._attributeCount {
		return "Unknown", errors.New("not enough data for classification")
	}
	unknownItem := Entity{"", attrValues[:model._attributeCount]}
	distances := make([]struct {
		string
		float64
	}, len(model._set))
	for i, entry := range model._set {
		dist := model._metric(unknownItem, entry)
		distances[i] = struct {
			string
			float64
		}{entry.ClassName, dist}
	}
	sort.Slice(distances, func(i, j int) bool { return distances[i].float64 < distances[j].float64 })
	frequencies := make(map[string]int)
	for _, val := range distances[:model.K] {
		frequencies[val.string]++
		if frequencies[val.string] > model.K/2 {
			return val.string, nil
		}
	}
	return "Unknown", errors.New("item cannot be classified")
}

var EuclideanDist Metric = func(a, b Entity) float64 {
	var res float64
	for i, val := range a.AttributeValues {
		res += (val - b.AttributeValues[i]) * (val - b.AttributeValues[i])
	}
	return math.Sqrt(res)
}
