package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type Chromosome struct {
	Gen     string
	Fitness float32
}

const MIN_ASCII int32 = 32
const MAX_ASCII int32 = 126

func randomChar() byte {
	return byte(rand.Int31n(MAX_ASCII-MIN_ASCII+1) + MIN_ASCII)
}

func rank(population []Chromosome) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})
}

// Genetic is a function to create genetic representation of data
func GenerateGen(length int) string {
	var gen strings.Builder
	var chars []byte

	for i := 0; i < length; i++ {
		chars = append(chars, randomChar())
	}
	gen.Write(chars)

	return gen.String()
}

// Fitness for calculate fitness value from compared data
func CalcFitness(genetic string, target string) float32 {
	targets := strings.Split(target, "")
	genetics := strings.Split(genetic, "")

	var correct int8
	for i := 0; i < len(target); i++ {
		var cmp int8
		if targets[i] == genetics[i] {
			cmp = 1
		}
		correct += cmp
	}

	return (float32(correct) / float32(len(target))) * 100.0
}

// CreatePopulation will generate population from given target and size
func CreatePopulation(target string, size int) []Chromosome {
	var population []Chromosome
	for i := 0; i < size; i++ {
		gen := GenerateGen(len(target))
		fitness := CalcFitness(gen, target)
		chromosome := Chromosome{
			Gen:     gen,
			Fitness: fitness,
		}
		population = append(population, chromosome)
	}

	return population
}

// Selection
func Selection(population []Chromosome) (Chromosome, Chromosome) {
	rank(population)
	return population[0], population[1]
}

// Crossover
func Crossover(first Chromosome, second Chromosome) (Chromosome, Chromosome) {
	var firstChild string = first.Gen
	var secondChild string = second.Gen

	point := len(first.Gen) / 2
	firstChild = strings.Replace(firstChild, firstChild[:point], second.Gen[:point], -1)
	secondChild = strings.Replace(secondChild, secondChild[:point], first.Gen[:point], -1)
	first.Gen = firstChild
	second.Gen = secondChild

	return first, second
}

// Mutate will change genetic in the Chromosome type.
// It will return copy of Chromosome
func Mutate(in Chromosome, rate float32) Chromosome {
	var out Chromosome = in
	for _, val := range out.Gen {
		if rand.Float32() <= rate {
			r := strings.Replace(out.Gen, string(val), string(randomChar()), -1)
			out.Gen = r
		}
	}

	return out
}

func Regeneration(ch []Chromosome, p []Chromosome) []Chromosome {
	rePopulate := make([]Chromosome, len(p))
	copy(rePopulate, p)

	rePopulate[len(rePopulate)-1] = ch[1]
	rePopulate[len(rePopulate)-2] = ch[0]

	return rePopulate
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var target string = "nada"

	var populationSize int = 100
	var mutationRate float32 = 0.3
	population := CreatePopulation(target, populationSize)
	isLoop := true
	i := 0
	for isLoop {
		firstParent, secondParent := Selection(population)

		firstChild, secondChild := Crossover(firstParent, secondParent)
		firstMutant := Mutate(firstChild, mutationRate)
		secondMutant := Mutate(secondChild, mutationRate)
		firstMutant.Fitness = CalcFitness(firstMutant.Gen, target)
		secondMutant.Fitness = CalcFitness(secondMutant.Gen, target)

		children := []Chromosome{firstMutant, secondMutant}
		population = Regeneration(children, population)

		if first, _ := Selection(population); first.Fitness == 100 {
			fmt.Printf("Iteration : %d\n", i)
			fmt.Printf("Gen       : %s\n", first.Gen)
			fmt.Printf("Fitness   : %.3f\n", first.Fitness)
			isLoop = false
		} else {
			fmt.Printf("Iteration : %d\n", i)
			fmt.Printf("Gen       : %s\n", first.Gen)
			fmt.Printf("Fitness   : %.3f\n", first.Fitness)
			clear()
			i++
			isLoop = true
		}
	}
}
