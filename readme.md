# Data Structures and Algorithms

## Linked List in Go

- What is a linked list?
- Why use them?
- Array vs Liked List?
- How it works?
- How to create a liked list

### Definitions

- A linier data structure, in which elements are <u>not</u> stored at contiguous memory locations.
- Array store elements in contiguous memory locations.
- Array good for modifying a certain element.
- Linked List good for inserting or deleting an element.
- Each element in a linked list is a node.
- Node is composed two parts,
  - the data that you want to store, can be really any type of data.
  - a pointer or reference to the next node in the list.
- To add new element,we create a new node which has a reference to the next node in the list.
- If we delete the node at index 1, we still need to make the pointer of the <u>head</u> node point to the node at index 2.
