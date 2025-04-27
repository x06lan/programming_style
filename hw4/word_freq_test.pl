:- use_module(library(plunit)).
:- use_module(word_freq).

:- begin_tests(word_freq).

test(filter_chars_and_normalize) :-
    filter_chars_and_normalize("This is a test, only a test.", Result),
    assertion(Result == "this is a test only a test ").

test(scan) :-
    scan("this is a test only a test ", Words),
    assertion(Words == ["this", "is", "a", "test", "only", "a", "test"]).

test(remove_stop_words) :-
    remove_stop_words(["this","is","a","test","only","a","test"], ["a","is","this"], Result),
    assertion(Result == ["test", "only", "test"]).

test(frequencies) :-
    frequencies(["test", "only", "test"], Freq),
    assertion(Freq == ["only"-1 ,"test"-2]).

test(sorted) :-
    sorted(["only"-1, "test"-2], Sorted),
    assertion(Sorted == ["test"-2, "only"-1]).

test(read_stop_words) :-
    read_stop_words('tmp_stop_words.txt', StopWords),
    assertion(StopWords==["a","is","this"]).

test(word_frequencies_output) :-
    once(word_frequencies('input.txt', 'tmp_stop_words.txt')).

:- end_tests(word_freq).

:- run_tests.