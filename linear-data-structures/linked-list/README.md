## Linked List  

A **Linked List** is a collection of nodes, where each node holds data and a reference (or pointer) to the next node. Unlike arrays, linked lists do not require contiguous memory allocation, making insertion and deletion operations more efficient.

### Types of Linked Lists  
Linked lists can be implemented in multiple ways based on how the nodes are connected:

- **Singly Linked List**: Each node points to the next node in the sequence.
- **Doubly Linked List**: Each node contains references to both the next and previous nodes.
- **Circular Linked List**: The last node points back to the first node, forming a circular connection.
- **Doubly Circular Linked List**: A combination of doubly and circular linked lists, where the last node connects to the first, and each node has references to both the next and previous nodes.

### Advantages of Linked Lists
- **Dynamic Size**: Can grow or shrink as needed.
- **Efficient Insertions/Deletions**: No need for shifting elements like in arrays.
- **Better Memory Utilization**: Uses memory as needed, avoiding preallocation.

### Disadvantages of Linked Lists
- **More Memory Usage**: Extra pointers require additional storage.
- **Slower Access**: Requires traversal from the head node to access elements (no direct indexing).