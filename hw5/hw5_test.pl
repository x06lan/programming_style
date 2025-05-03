:- use_module(library(plunit)).
:- use_module(hw5).

:- begin_tests(hw5).

test(child_success) :-
    children(s1, C),
    C = [l4, l3].

test(descendants_success) :-
    descendants(s1, D),
    D = [s2, l4, l3].

test(new_instance_stage) :-
    new_instance(i1, stage).

test(new_instance_lane_fail, [fail]) :-
    new_instance(i1, lane).

test(valid_children_check) :-
    valid_children(s1).

test(valid_children_check,[fail]) :-
    valid_children(i1).

test(add_child, [true]) :-
    add_child(l3, s1).

test(add_child_duplicate_parent, [true]) :-
    \+ add_child(l5, s2).  % l5 already has a parent (l4)

test(add_child_ancestor_cycle, [true]) :-
    \+ add_child(s1, s2).  % s1 is an ancestor of s2

test(add_child_descendant_cycle, [true]) :-
    \+ add_child(s2, s1).  % s2 is a descendant of s1

test(add_child_class_mismatch, [true]) :-
    \+ add_child(stage_instance, lane_instance).  % hypothetical mismatched concrete classes


:- end_tests(hw5).

:- run_tests.