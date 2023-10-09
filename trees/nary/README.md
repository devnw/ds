# Generic N-Ary Trees

This is a generic implementation of an N-Ary tree. There are no limitations
on the number of children a node can have. The tree is implemented using
a slice of pointers to the children nodes.

Currently no struct fields are exported so that internal implementation can
be changed without breaking the API. Exported functions are provided to
manipulate the tree.
