# Fenwick Tree

A Fenwick tree is a data structure that holds an ordered collection and supports
the operations `sum` and `update`, both in *O(log n)* time.

In its most basic form the tree stores an array of integers. `Sum` calculates
the cumulative total of the first *n* integers and `update` modifies an element. Theoretically one is not limited to addition of integers when using
Fenwick trees. See [this](http://tobin.cc/blog/fenwick/) blog post for more information.

## Directory Listing

**fenwick-add/** Fenwick tree using addition for the update operation.  
**fenwick-mul/** Fenwick tree using multiplication for the update operation.  
**fenwick-get-set/**
Fenwick tree using addition for the update operation. Also implementing Getter
and Setter functions.
