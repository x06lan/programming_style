# HW5 (Deadline: 5/5 Mon. 23:59)

Due May 5, 2025 11:59 PM
•
Closes May 5, 2025 11:59 PM
Instructions
The revised requirements are marked as this type.
In your unit tests for the last goal: valid_children, you can just test the true condition.

Write a Prolog program for the requirements subject to the implementation constraints:
Classes, relationships, and instance are presented as Prolog facts or rules.

There are three classes: lane, stage, and swimlane.
Lane is abstract.
Stage and swimlane are concrete subclasses of lane.
Only a concrete class can generate an instance.
An instance can have zero or more children (sub-lanes) attached to it.
All children of a lane must be of the same concrete class.
An instance cannot be a child of any descendant.
An instance cannot be a child of itself.
An instance that is a child of instance x cannot be the child of instance y.  

![](https://imgur.com/vu8dk7U.png)



The Prolog program can accept the following goals with the results that follow:
```
?- new_instance(i, stage)
true
?- new_instance(i, lane)
false
% assuming s1, s2 are instances of stages and l3, l4, and l5 are instances of swimlane
?- add_child(l3, s1). % adding l3 as a child a s1
true
?- children(s1, L).
L = [l3]
?- descendants(s1, L).
L = [l3]
?- add_child(s2, s1).
Child not added: cannot add a child with different class
true.
?- children(s1, L).
L = [l3]
?- descendants(s1, L).
L = [l3]
?- add_child(l4, s1). % adding l4 as a child a s1
true
?- add_child(s2, l3). % adding s2 as a child a l3
true
?- children(s1, L).
L = [l3, l4].
?- descendants(s1, L).
L = [l3, l4, s2].
?- add_child(s1, s2).
Child not added: cannot add an ancestor as a new child
true.
?- add_child(s2, s1).
Child not added: cannot add a descendant as a new child
true.
?- add_child(l5,l4).
true .
?- add_child(l5,s2).
Child not added: cannot add a child that is already a child of another parent
true.
% valid_children(instance) checks if instance satisfies the "children are of same concrete class" rule.
?- valid_children(s1).
true.
?- valid_children(l3).
true.
```


## Folder Structure
```
學號_姓名_hw5/
├── hw5.pl
└── hw5_test.pl
```

## Grading Criteria
- Source Code + Test code: (50%)
- Tests from TA: (50%)