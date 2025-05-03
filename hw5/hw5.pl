subclass(stage, lane).
subclass(swimlane, lane).

concrete(stage).
concrete(swimlane).


new_instance(_, Class) :-
    concrete(Class).

children(Parent, ChildrenList) :-
    findall(Child, add_child(Child, Parent), ChildrenList).

descendants(Parent, AllDescendants) :-
    descendants_acc(Parent, [], AllDescendants).

descendants_acc(Node, Visited, Descendants) :-
    children(Node, Children),
    include(\=(Visited), Children, NewChildren),
    append(NewChildren, Visited, UpdatedVisited),
    foldl(descendants_acc, NewChildren, UpdatedVisited, Descendants).

valid_children(Instance) :-
    instance(Instance, Class),
    concrete(Class).

/*
new_instance(+Class, +Instance) is det.
children(+Parent, -Children) is det.
descendant(+Parent, -Descendants) is det.
valid_child(+Child) is det.
*/


instance(i1, lane).

instance(s1,stage).
instance(s2,stage).

instance(l3,swimlane).
instance(l4,swimlane).
instance(l5,swimlane).

add_child(l4, s1).
add_child(l3, s1).
add_child(s2, l3).
