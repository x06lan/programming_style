## HW3: Implementing the Template Method Pattern in Go  

### **Deadline**: 4/7 (Mon.) 23:59  

### **Background**
In software development, many applications need to generate documents in different formats, such as plain text documents and HTML documents. The document generation process typically follows these common steps:

1. Prepare the raw data
2. Format the content
3. Save the formatted content

Although these steps are fixed in sequence, different document types may have variations in implementation. For example:
- HTML documents require tags like `<html>` and `<body>`.
- Plain text documents may only need simple strings.

To unify this process, we can use the **Template Method Pattern**, which encapsulates these steps within a **Base Document Generator** while allowing different document types to define their specific behaviors.

---

## **Assignment Requirements**

Based on the above concept, design a **Document Generator Framework** and complete the following tasks.

### **1. Define the DocumentGenerator Interface**
The `DocumentGenerator` interface defines the fundamental steps for document generation. Each document type must implement the following methods:

```go
PrepareData() string      // Prepare the raw content of the document
FormatContent(data string) string  // Format the content
Save(content string) string  // Save the formatted content
```
Each step’s implementation will depend on the specific document type.

---

### **2. Implement the BaseGenerator Class**
Design a **Base Document Generator (`BaseGenerator`)** that contains the `Generate()` method, which executes the three steps defined in the `DocumentGenerator` interface sequentially. The `Generate()` method does not implement specific logic; instead, it relies on the document type’s implementation of the `DocumentGenerator` interface.

---

### **3. Implement TextDocument and HTMLDocument Classes**
Implement the following two document types and ensure they conform to the `DocumentGenerator` interface:

#### **(1) Plain Text Document (`TextDocument`)**
- `PrepareData()`: Returns `"This is the raw text data."`
- `FormatContent(data string)`: Returns `"Formatted Text: {data}"`
- `Save(content string)`: Returns `"Saving text document: {formatted_content}"`

##### **Expected Output Example:**
```
Saving text document: Formatted Text: This is the raw text data.
```

#### **(2) HTML Document (`HTMLDocument`)**
- `PrepareData()`: Returns `"<html><body>This is raw HTML data.</body></html>"`
- `FormatContent(data string)`: Returns `"<div>{data}</div>"`
- `Save(content string)`: Returns `"Saving HTML document: {formatted_content}"`

##### **Expected Output Example:**
```
Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>
```

---

### **4. Write Unit Tests**
Write test cases for `TextDocument` and `HTMLDocument` to ensure the `Generate()` method executes correctly.

---

## **Grading Criteria**
- **Correct implementation of `TextDocument` and `HTMLDocument` (40%)**
- **Generate() method follows Template Method Pattern (30%)**
- **Completeness and correctness of test cases (30%)**


## How to run
```
cd hw3
go test
```