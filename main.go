package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
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

func Selection(population []Chromosome) (Chromosome, Chromosome) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})

	return population[0], population[1]
}

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

func main() {
	rand.Seed(time.Now().UnixNano())

	var target string
	target = "wiro sableng"

	size := 10
	population := CreatePopulation(target, size)
	firstParent, secondParent := Selection(population)

	firstChild, secondChild := Crossover(firstParent, secondParent)
	// Mutation
	var mutationRate float32 = 0.5
	firstMutant := Mutate(firstChild, mutationRate)
	secondMutant := Mutate(secondChild, mutationRate)

	fmt.Println("Parent : ")
	fmt.Println(firstParent)
	fmt.Println(secondParent)

	fmt.Println("Child : ")
	fmt.Println(firstChild)
	fmt.Println(secondChild)

	fmt.Println("Mutation : ")
	fmt.Println(firstMutant)
	fmt.Println(secondMutant)
}
