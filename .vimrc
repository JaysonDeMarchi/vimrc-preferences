syntax on
set tabstop=4 shiftwidth=4 autoindent expandtab
colorscheme desert

" -- INSERT MODE --

" closing brackets
inoremap (<CR> (<CR><TAB><CR>)<ESC><<<UP>$a
inoremap (;<CR> (<CR><TAB><CR>);<ESC><<<UP>$a
inoremap [<CR> [<CR><TAB><CR>]<ESC><<<UP>$a
inoremap [;<CR> [<CR><TAB><CR>];<ESC><<<UP>$a
inoremap {<CR> {<CR><TAB><CR>}<ESC><<<UP>$a
inoremap {;<CR> {<CR><TAB><CR>};<ESC><<<UP>$a

" closing block quotes
inoremap /**<SPACE> /**<SPACE><SPACE>*/<LEFT><LEFT><LEFT>
inoremap /**<CR> /**<CR><SPACE>*<CR>*/<UP><SPACE>
