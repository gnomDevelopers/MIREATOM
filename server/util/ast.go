package util

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Node представляет узел AST
type Node struct {
	Type     string  // Тип узла (например, "Plus", "Times", "Variable")
	Value    string  // Значение узла, если применимо
	Children []*Node // Дочерние узлы
}

// Formula представляет структуру с исходной формулой и AST
type Formula struct {
	Original string
	AST      *Node
}

// Слайс для хранения формул
var formulas []Formula

// Парсер для построения AST из LaTeX-формулы
func parseFormula(formula string) (*Node, error) {
	formula = strings.TrimSpace(formula)

	// Проверяем на операторы от высокого к низкому приоритету
	operators := []string{"=", "+", "-", "*", "/", "^"}
	for _, operator := range operators {
		parts := splitByOperator(formula, operator)
		if len(parts) > 1 {
			children := []*Node{}
			for _, part := range parts {
				child, err := parseFormula(part)
				if err != nil {
					return nil, err
				}
				children = append(children, child)
			}

			// Для коммутативных операций сортируем дочерние узлы
			if operator == "+" || operator == "*" {
				sort.Slice(children, func(i, j int) bool {
					return nodeToString(children[i]) < nodeToString(children[j])
				})
			}

			return &Node{
				Type:     operator,
				Children: children,
			}, nil
		}
	}

	// Проверяем на скобки и убираем их
	if strings.HasPrefix(formula, "(") && strings.HasSuffix(formula, ")") {
		return parseFormula(formula[1 : len(formula)-1])
	}

	// Если ничего не подходит, это переменная или значение
	return &Node{Type: "Variable", Value: formula}, nil
}

// Разделение строки по оператору с учетом структуры
func splitByOperator(formula, operator string) []string {
	parts := []string{}
	current := ""
	balance := 0
	for _, char := range formula {
		switch char {
		case '(':
			balance++
		case ')':
			balance--
		}
		if string(char) == operator && balance == 0 {
			parts = append(parts, strings.TrimSpace(current))
			current = ""
		} else {
			current += string(char)
		}
	}
	if current != "" {
		parts = append(parts, strings.TrimSpace(current))
	}
	return parts
}

// Сравнение двух деревьев AST
func areTreesEqual(a, b *Node) bool {
	if a.Type != b.Type || a.Value != b.Value {
		return false
	}
	if len(a.Children) != len(b.Children) {
		return false
	}
	for i := range a.Children {
		if !areTreesEqual(a.Children[i], b.Children[i]) {
			return false
		}
	}
	return true
}

// Поиск всех возможных общих поддеревьев между двумя AST
func findAllCommonSubtrees(a, b *Node, visited map[string]bool) []*Node {
	if a == nil || b == nil {
		return nil
	}

	var commonSubtrees []*Node
	if a.Type == b.Type && a.Value == b.Value {
		common := &Node{Type: a.Type, Value: a.Value}
		for i := 0; i < len(a.Children) && i < len(b.Children); i++ {
			childCommon := findAllCommonSubtrees(a.Children[i], b.Children[i], visited)
			if len(childCommon) > 0 {
				common.Children = append(common.Children, childCommon...)
			}
		}
		key := nodeToString(common)
		if !visited[key] {
			visited[key] = true
			commonSubtrees = append(commonSubtrees, common)
		}
	}

	// Рекурсивный поиск на уровне потомков
	for _, childA := range a.Children {
		commonSubtrees = append(commonSubtrees, findAllCommonSubtrees(childA, b, visited)...)
	}
	for _, childB := range b.Children {
		commonSubtrees = append(commonSubtrees, findAllCommonSubtrees(a, childB, visited)...)
	}

	return commonSubtrees
}

// Преобразование дерева в строку для сортировки
func nodeToString(node *Node) string {
	if node == nil {
		return ""
	}
	if len(node.Children) == 0 {
		return node.Value
	}
	childrenStrings := []string{}
	for _, child := range node.Children {
		childrenStrings = append(childrenStrings, nodeToString(child))
	}
	sort.Strings(childrenStrings)
	return fmt.Sprintf("%s(%s)", node.Type, strings.Join(childrenStrings, ","))
}

// Проверка, есть ли совпадение среди сохраненных деревьев
func checkSimilarity(newTree *Node) ([]*Node, bool) {
	for _, formula := range formulas {
		if areTreesEqual(formula.AST, newTree) {
			return nil, true
		}
		visited := make(map[string]bool)
		commonSubtrees := findAllCommonSubtrees(formula.AST, newTree, visited)
		if len(commonSubtrees) > 0 {
			return commonSubtrees, false
		}
	}
	return nil, false
}

// Добавление формулы в список, если она уникальна
func addFormula(formula string) error {
	// Удаляем внешний формат LaTeX
	latexRegex := regexp.MustCompile(`\\\[(.*?)\\\]`)
	matches := latexRegex.FindStringSubmatch(formula)
	if len(matches) < 2 {
		return fmt.Errorf("invalid LaTeX formula")
	}
	formulaContent := matches[1]
	tree, err := parseFormula(formulaContent)
	if err != nil {
		return fmt.Errorf("failed to parse formula: %v", err)
	}

	commonSubtrees, isDuplicate := checkSimilarity(tree)
	if isDuplicate {
		return fmt.Errorf("formula already exists")
	}
	if len(commonSubtrees) > 0 {
		fmt.Println("Similar parts found:")
		for _, common := range commonSubtrees {
			printTree(common, "  ")
		}
	}

	formulas = append(formulas, Formula{
		Original: formula,
		AST:      tree,
	})
	return nil
}

// Печать AST для отладки
func printTree(node *Node, prefix string) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s (%s)\n", prefix, node.Type, node.Value)
	for _, child := range node.Children {
		printTree(child, prefix+"  ")
	}
}

// Основная функция
func main() {
	// Примеры ввода формул
	inputs := []string{
		"\\[A = \\pi r^{2}\\]",
		"\\[B = \\pi r^{2}\\]",
		"\\[x + \\frac{2}{5}\\]",
		"\\[x + y + z\\]",
		"\\[z + x + y\\]", // Перестановка
		"\\[x + y\\]",
		"\\[A = \\pi r^{2}\\]", // Дубликат
	}

	for _, formula := range inputs {
		err := addFormula(formula)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Added formula: %s\n", formula)
		}
	}

	fmt.Println("Stored formulas:")
	for _, formula := range formulas {
		fmt.Printf("Original: %s\n", formula.Original)
		printTree(formula.AST, "  ")
	}
}
