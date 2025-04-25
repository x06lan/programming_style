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
    remove_stop_words(["this","is","a","test","only","a","test"], ["a","is","the"], Result),
    assertion(Result == ["this", "test", "only", "test"]).

test(frequencies) :-
    frequencies(["test", "only", "test"], Freq),
    assertion(Freq == ["only"-1 ,"test"-2]).

test(sorted) :-
    sorted(["only"-1, "test"-2], Sorted),
    assertion(Sorted == ["test"-2, "only"-1]).

test(read_stop_words, [setup(open('tmp_stop_words.txt', write, S)), cleanup(delete_file('tmp_stop_words.txt'))]) :-
    write(S, "a,is,the"), close(S),
    read_stop_words('tmp_stop_words.txt', StopWords),
    assertion(StopWords==["a","is","the"]).

test(word_frequencies_output, [setup((open('input.txt', write, S1), write(S1, "This is a test, only a test."), close(S1),
                                      open('tmp_stop_words.txt', write, S2), write(S2, "a,is,the"), close(S2))),
                               cleanup((delete_file('input.txt'), delete_file('tmp_stop_words.txt')))]) :-
    once(word_frequencies('input.txt', 'tmp_stop_words.txt')).

:- end_tests(word_freq).

:- run_tests.