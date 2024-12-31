# GoAlgoChain

**GoAlgoChain** is a powerful and flexible data processing pipeline library for Go. It offers a fluent API for chaining
various data operations and algorithms, enabling developers to build efficient and readable data processing workflows
with ease.

[![Go Reference](https://pkg.go.dev/badge/github.com/NaokiOouchi/GoAlgoChain/pkg/algo.svg)](https://pkg.go.dev/github.com/NaokiOouchi/GoAlgoChain/pkg/algo)
[![Go Version](https://img.shields.io/github/go-mod/go-version/NaokiOouchi/GoAlgoChain)](https://github.com/NaokiOouchi/GoAlgoChain/)
[![Github license](https://img.shields.io/github/license/NaokiOouchi/GoAlgoChain)](https://github.com/NaokiOouchi/GoAlgoChain/blob/main/LICENSE)
[![Github issues](https://img.shields.io/github/issues/NaokiOouchi/GoAlgoChain)](https://github.com/NaokiOouchi/GoAlgoChain/issues)
[![Github forks](https://img.shields.io/github/forks/NaokiOouchi/GoAlgoChain)](https://github.com/NaokiOouchi/GoAlgoChain/network/members)
[![Github stars](https://img.shields.io/github/stars/NaokiOouchi/GoAlgoChain)](https://github.com/NaokiOouchi/GoAlgoChain/stargazers)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Benchmark Results](#benchmark-results)
- [Documentation](#documentation)
- [License](#license)
- [Contact](#contact)

## Features

- **Fluent API**: Chain multiple operations seamlessly for clear and concise data processing.
- **Generic Support**: Utilize Go's generics to handle various data types with type safety.
- **Comprehensive Operations**:
    - **Filtering**: `Filter`, `Distinct`
    - **Searching**: `BinarySearch`, `LinearSearch`, `Find`
    - **Sorting**: `QuickSort`, `MergeSort`, `HeapSort`
    - **Transforming**: `Map`, `Reduce`, `GroupBy`, `Take`, `Skip`
- **Performance Optimized**: Implemented with performance and memory efficiency in mind.
- **Extensible**: Easily add custom operations to extend functionality.

## Installation

Install GoAlgoChain using `go get`:

```bash
go get github.com/NaokiOouchi/GoAlgoChain/pkg/algo
```

## Usage

### Basic Pipeline Operations

```go
import "github.com/NaokiOouchi/GoAlgoChain/pkg/algo"
// Create a pipeline with initial data
numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
result, err := algo.NewPipelineWithData(numbers).
    Filter(func (x int) bool { return x > 2 }).
    Map(func (x int) int { return x * 2 }).
    QuickSort(func (a, b int) bool { return a < b }).
    Execute()

// Alternative initialization
pipeline := algo.NewPipeline[int]().
    WithData(numbers)
```

### Sorting Operations

```go
// QuickSort
sorted, _ := algo.NewPipelineWithData(items).
    QuickSort(func(a, b Item) bool { return a.ID < b.ID }).
    Execute()

// MergeSort
sorted, _ := algo.NewPipelineWithData(items).
    MergeSort(func(a, b Item) bool { return a.Price < b.Price }).
    Execute()

// HeapSort
sorted, _ := algo.NewPipelineWithData(items).
    HeapSort(func(a, b Item) bool { return a.Priority > b.Priority }).
    Execute()
```

### Searching Operations
```go
// Binary Search (requires sorted data)
found, _ := algo.NewPipelineWithData(sortedItems).
    BinarySearch(func(item Item) bool { return item.ID == targetID }).
    Execute()

// Linear Search
found, _ := algo.NewPipelineWithData(items).
    LinearSearch(func(item Item) bool { return item.Name == targetName }).
    Execute()
// Find
found, _ := algo.NewPipelineWithData(items).
    Find(func(item Item) bool { return item.ID == targetID }).
    Execute()
```

### Transform Operations
```go
// Filter
filtered, _ := algo.NewPipelineWithData(items).
    Filter(func(item Item) bool { return item.Price > 100 }).
    Execute()

// Map
transformed, _ := algo.NewPipelineWithData(items).
    Map(func(item Item) Item { 
        item.Price *= 1.1  // Add 10% to price
        return item 
    }).
    Execute()

// Reduce
sum, _ := algo.NewPipelineWithData(items).
    Reduce(func(acc, item Item) Item {
        acc.Total += item.Amount
        return acc
    }).
    Execute()
```

### Grouping and Distinct Operations
```go
// GroupBy
grouped := algo.GroupBy(items, func(item Item) string { 
    return item.Category 
})

// Distinct
unique, _ := algo.NewPipelineWithData(items).
    Distinct(func(a, b Item) bool { return a.ID == b.ID }).
    Execute()
```

### Pagination Operations
```go
// Skip and Take for pagination
page, _ := algo.NewPipelineWithData(items).
    Skip(10).  // Skip first 10 items
    Take(5).   // Take next 5 items
    Execute()
```

### Complex Pipelines Example
```go
result, _ := algo.NewPipelineWithData(orders).
    Filter(func(o Order) bool { return o.Status == "completed" }).
    Distinct(func(a, b Order) bool { return a.UserID == b.UserID }).
    QuickSort(func(a, b Order) bool { return a.Amount > b.Amount }).
    Take(5).
    Execute()
```

## Advanced Examples
Explore more advanced usage scenarios in the [Examples](/example) section below.

# Benchmark
## Analysis
- Highly Efficient Operations:
  - BinarySearchOperation, LinearSearch, Reduce, Skip, and Take demonstrate exceptional performance with minimal memory usage and allocations.
- Performance Bottlenecks:
  - Operations like FilterOperation, Find, GroupBy, HeapSort, MergeSort, QuickSort, and Map show significant execution times and memory usage. Optimization efforts are needed for these operations to enhance performance and reduce memory footprint.

# Documentation
Comprehensive documentation is available through GoDoc. You can access it here:

[GoDoc Documentation](https://pkg.go.dev/github.com/NaokiOouchi/GoAlgoChain/pkg/algo)

# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Contact
For any questions or inquiries, please contact [X](https://x.com/NaoNoaNaoNoaN) .
