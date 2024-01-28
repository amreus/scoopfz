ACTION: DEFAULT BINDINGS (NOTES):

abort ctrl-c ctrl-g ctrl-q esc
accept enter double-click
backward-char ctrl-b left
backward-delete-char ctrl-h bspace
backward-kill-word alt-bs
backward-word alt-b shift-left
beginning-of-line ctrl-a home
clear-screen ctrl-l
delete-char del
delete-char/eof ctrl-d (same as delete-char except aborts fzf if query is empty)
down ctrl-j ctrl-n down
end-of-line ctrl-e end
forward-char ctrl-f right
forward-word alt-f shift-right
kill-word alt-d
next-history (ctrl-n on --history)
next-selected (move to the next selected item)
page-down pgdn
page-up pgup
half-page-down
half-page-up
hide-preview
offset-down (similar to CTRL-E of Vim)
offset-up (similar to CTRL-Y of Vim)
pos(...) (move cursor to the numeric position; negative number to count from the end)
prev-history (ctrl-p on --history)
prev-selected (move to the previous selected item)
preview(...) (see below for the details)
preview-down shift-down
preview-up shift-up
preview-page-down
preview-page-up
preview-half-page-down
preview-half-page-up
preview-bottom
preview-top
print-query (print query and exit)
put (put the character to the prompt)
put(...) (put the given string to the prompt)
refresh-preview
rebind(...) (rebind bindings after unbind)
reload(...) (see below for the details)
reload-sync(...) (see below for the details)
replace-query (replace query string with the current selection)
select
select-all (select all matches)
show-preview
toggle (right-click)
toggle-all (toggle all matches)
toggle+down ctrl-i (tab)
toggle-header
toggle-in (--layout=reverse* ? toggle+up : toggle+down)
toggle-out (--layout=reverse* ? toggle+down : toggle+up)
toggle-preview
toggle-preview-wrap
toggle-search (toggle search functionality)
toggle-sort
toggle-track
toggle+up btab (shift-tab)
track (track the current item; automatically disabled if focus changes)
transform-border-label(...) (transform border label using an external command)
transform-header(...) (transform header using an external command)
transform-preview-label(...) (transform preview label using an external command)
transform-prompt(...) (transform prompt string using an external command)
transform-query(...) (transform query string using an external command)
unbind(...) (unbind bindings)
unix-line-discard ctrl-u
unix-word-rubout ctrl-w
up ctrl-k ctrl-p up
yank ctrl-y
