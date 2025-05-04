subclass(stage, lane).
subclass(swimlane, lane).

concrete(stage).
concrete(swimlane).

:- dynamic instance/2.
:- dynamic add_child/2.
:- dynamic have_children/2.

new_instance(_, Class) :-
    concrete(Class),
    \+ instance(_, Class),
    assertz(instance(_, Class)).

new_instance(_, Class) :-
    concrete(Class),
    instance(_, Class).

children(Parent, ChildrenList) :-
    findall(Child, have_children(Child, Parent), ChildrenList).

descendants(Parent, AllDescendants) :-
    descendants_acc(Parent, [], AllDescendants).

descendants_acc(Node, Visited, Descendants) :-
    children(Node, Children),
    include(\=(Visited), Children, NewChildren),
    append(NewChildren, Visited, UpdatedVisited),
    foldl(descendants_acc, NewChildren, UpdatedVisited, Descendants).

% valid_children(instance) checks if instance satisfies the "children are of same concrete class" rule.
valid_children(Instance) :-
    instance(Instance, Class),
    concrete(Class).

add_child(Child, Parent) :-
    have_children(Child, Parent).

add_child(Child, Parent) :-
    \+have_children(Child, Parent),
    Child \= Parent,
    \+have_children(Child, _),
    \+have_children(_ , Child),
    assertz(have_children(Child, Parent)).

add_child(Child, Parent) :-
    instance(Child, _),
    instance(Parent, _).





/*
new_instance(+Class, +Instance) is det.
children(+Parent, -Children) is det.
descendant(+Parent, -Descendants) is det.
valid_child(+Child) is det.
*/


/*


:- new_instance(s1, stage).
:- new_instance(s2, stage).
:- new_instance(l3, swimlane).
:- new_instance(l4, swimlane).

:- add_child(l3, s1).
:- add_child(l4, s1).
:- add_child(s2, l3).



*/
