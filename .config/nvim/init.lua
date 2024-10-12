require("options")
require("keymaps")

require("config.lazy")

vim.opt.termguicolors = true
require("bufferline").setup{}

-- Lualine
require('lualine').setup()

-- Theme/ColorScheme
vim.cmd.colorscheme "catppuccin-macchiato"

-- testing configs separate...
-- IT WORKS!
-- keymaps doesn't work!
require("treesitter_config")

require("telescope_config")

require("bufferline_config")

require("lualine_config")

require("hop_config")

require("alpha_config")

require("lsp_config")

require("autopairs_config")

require("whichkey_config")

require("treesitter_config")

require("cmp_config")

require('lspconfig').clangd.setup({})
require('lspconfig').rust_analyzer.setup({})
-- You would add this setup function after calling lsp_zero.extend_lspconfig()