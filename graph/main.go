package main

import "fmt"

// const visited = make([]string, 0)
// const graph = make(map[string][]string)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func graph_dfs(graph_string string, visited []string, graph map[string][]string) {
	fmt.Println(graph_string)
	visited = append(visited, graph_string)
	for _, graph_label := range graph[graph_string] {
		if !contains(visited, graph_label) {
			graph_dfs(graph_label, visited, graph)
		}
	}
}

func graph_bfs(graph_string string, visited, queue []string, graph map[string][]string) {
	fmt.Println(graph_string)
	visited = append(visited, graph_string)
	for _, graph_label := range graph[graph_string] {
		if !contains(visited, graph_label) {
			graph_bfs(graph_label, visited, graph)
		}
	}
}

func main() {
	visited := make([]string, 0)
	graph := make(map[string][]string)
	r := make(map[string][]string)
	r["A"] = []string{"B", "C", "D"}
	r["B"] = []string{"A", "E"}
	r["C"] = []string{"A", "E"}
	r["D"] = []string{"A"}
	r["E"] = []string{"B", "C", "F"}
	r["F"] = []string{"E"}
}
