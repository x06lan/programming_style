:- use_module(library(assoc)).

read_file(FileName,String) :-
    open(FileName, read, Stream),
    read_string(Stream, _, String0),
    close(Stream),
    String = String0.

read_stop_words(Path,WordList) :-
    read_file(Path, String),
    split_string(String, "," , "", WordList0),
    WordList = WordList0.

filter_chars_and_normalize(Text, FilteredText) :-
    re_replace("[\\W_]+"/g, " ", Text, FilteredText0),
    string_lower(FilteredText0, FilteredText1),
    FilteredText = FilteredText1.

scan(String,WordsList) :-
    split_string(String, " " , "\n", WordList0),
    exclude(==( ""), WordList0, WordsList).
    
is_stop_word(StopWords,Word) :-
    member(Word, StopWords).

remove_stop_words(WordList, StopWords, FilteredWordList) :-
    exclude(is_stop_word(StopWords), WordList, FilteredWordList).

frequencies(Words, WordFreqList) :-
    empty_assoc(Empty),
    frequencies_assoc(Words, Empty, Assoc),
    assoc_to_list(Assoc, WordFreqList).

frequencies_assoc([], Assoc, Assoc).
frequencies_assoc([Word | Rest], Acc, Result) :-
    ( get_assoc(Word, Acc, Count) ->
        NewCount is Count + 1,
        put_assoc(Word, Acc, NewCount, NewAcc)
    ;
        put_assoc(Word, Acc, 1, NewAcc)
    ),
    frequencies_assoc(Rest, NewAcc, Result).


compare_desc(Delta, _-V1, _-V2) :-
    compare(Delta1, V2, V1),  % flip V1 and V2 to get descending
    Delta = Delta1.

sorted(WordList, SortedWordList) :-
    predsort(compare_desc, WordList, SortedWordList).

print_first([], _).
print_first(_, 0).
print_first([Word-Freq | T], Count) :-
    Count > 0,
    format("~w - ~d~n", [Word,Freq]),
    Count1 is Count - 1,
    print_first(T, Count1).

word_frequencies(File, StopWordsFile):-
    read_file(File, String),
    filter_chars_and_normalize(String, FilteredString),
    scan(FilteredString, WordList),
    read_stop_words(StopWordsFile, StopWords),
    remove_stop_words(WordList, StopWords, FilteredWordList),
    frequencies(FilteredWordList, WordFreq),
    sorted(WordFreq, SortedWordFreq),
    print_first(SortedWordFreq,25).

/*
:- word_frequencies("pride-and-prejudice.txt", 'stop_words.txt').
*/