:- use_module(library(plunit)).
:- use_module(hw5).


:- begin_tests(hw5).

test(new_instance_concrete) :-
    once(new_instance(s1, stage)),
    instance(s1, stage).

test(new_instance_non_concrete, [fail]) :-
    new_instance(i,lane).  % lane is not concrete

test(add_valid_child) :-
    once(new_instance(s1, stage)),
    once(new_instance(l3, swimlane)),
    once(add_child(l3, s1)),
    children(s1, L),
    assertion(L == [l3]).
    
test(add_descendants_child) :-
    once(new_instance(s1, stage)),
    once(new_instance(l3, swimlane)),
    once(add_child(l3, s1)),
    descendants(s1, L),
    assertion(L == [l3]).

test(add_additional_child) :-
    once(new_instance(s2, stage)),
    once(new_instance(l3, swimlane)),
    once(new_instance(l4, swimlane)),
    once(add_child(l4, s1)),
    once(add_child(s2, l3)),
    children(s1, D),
    sort(D, Sorted),
    assertion(Sorted == [l3, l4]).

test(descendants_simple) :-
    once(add_child(s2, l3)),
    descendants(s1, D),
    sort(D, Sorted),
    assertion(Sorted == [l3, l4, s2]).

test(add_different_class_fails) :-
    once(new_instance(s2,stage)),
    once(add_child(s2, s1)).

test(add_different_class_fails) :-
    once(new_instance(s2,stage)),
    once(add_child(s1, s2)).

test(add_different_class) :-
    once(new_instance(l5, swimlane)),
    once(add_child(l5, l4)).

test(add_different_class) :-
    once(new_instance(l5, swimlane)),
    once(add_child(l5, s2)).

test(valid_children) :-
    once(valid_children(s1)).
test(valid_children) :-
    once(valid_children(l3)).


:- end_tests(hw5).

:- run_tests.