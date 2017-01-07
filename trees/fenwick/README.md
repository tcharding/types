# Fenwick Tree

A Fenwick tree is a data structure that holds an ordered collection of
numbers. It supports two operations, `Sum` and `Update`.

Fenwick trees can be implemented to use any binary operation to update the
values in the data structure.

`Sum` sums the first n numbers in the collection. `Update` modifies the element at a given index using
whichever binary operation the implementation uses.

## Directory Listing

**fenwick-add/** Fenwick tree using addition for the update operation.  
**fenwick-mul/** Fenwick tree using multiplication for the update operation.  
**fenwick-get-set/**
Fenwick tree using addition for the update operation. Also implementing Getter
and Setter functions.
