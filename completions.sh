#!/usr/bin/env bash

# This also works for zsh: https://zsh.sourceforge.io/Doc/Release/Completion-System.html#Completion-System
_main()
{
    COMPREPLY=()

    local subcommands="completions.create messages.create messages.count_tokens messages.batches.create messages.batches.retrieve messages.batches.list messages.batches.delete messages.batches.cancel models.retrieve models.list"

    if [[ "$COMP_CWORD" -eq 1 ]]
    then
      local cur="${COMP_WORDS[COMP_CWORD]}"
      COMPREPLY=( $(compgen -W "$subcommands" -- "$cur") )
      return
    fi

    local subcommand="${COMP_WORDS[1]}"
    local flags
    case "$subcommand" in
      completions.create)
        flags="--max-tokens-to-sample --model --prompt --metadata.user_id --stop-sequences --+stop_sequence --temperature --top-k --top-p"
        ;;
      messages.create)
        flags="--max-tokens --messages.content.text --messages.content.type --messages.content.cache_control.type --messages.content.source.data --messages.content.source.media_type --messages.content.source.type --messages.content.id --messages.content.name --messages.content.tool_use_id --messages.content.content.text --messages.content.content.type --messages.content.content.cache_control.type --messages.content.content.source.data --messages.content.content.source.media_type --messages.content.content.source.type --messages.content.+content --messages.content.is_error --messages.+content --messages.role --+message --model --metadata.user_id --stop-sequences --+stop_sequence --system.text --system.type --system.cache_control.type --+system --temperature --tool-choice.type --tool-choice.disable_parallel_tool_use --tool-choice.name --tools.name --tools.cache_control.type --tools.description --+tool --top-k --top-p"
        ;;
      messages.count_tokens)
        flags="--messages.content.text --messages.content.type --messages.content.cache_control.type --messages.content.source.data --messages.content.source.media_type --messages.content.source.type --messages.content.id --messages.content.name --messages.content.tool_use_id --messages.content.content.text --messages.content.content.type --messages.content.content.cache_control.type --messages.content.content.source.data --messages.content.content.source.media_type --messages.content.content.source.type --messages.content.+content --messages.content.is_error --messages.+content --messages.role --+message --model --system --system.text --system.type --system.cache_control.type --+system --tool-choice.type --tool-choice.disable_parallel_tool_use --tool-choice.name --tools.name --tools.cache_control.type --tools.description --+tool"
        ;;
      messages.batches.create)
        flags="--requests.custom_id --requests.params.max_tokens --requests.params.messages.content.text --requests.params.messages.content.type --requests.params.messages.content.cache_control.type --requests.params.messages.content.source.data --requests.params.messages.content.source.media_type --requests.params.messages.content.source.type --requests.params.messages.content.id --requests.params.messages.content.name --requests.params.messages.content.tool_use_id --requests.params.messages.content.content.text --requests.params.messages.content.content.type --requests.params.messages.content.content.cache_control.type --requests.params.messages.content.content.source.data --requests.params.messages.content.content.source.media_type --requests.params.messages.content.content.source.type --requests.params.messages.content.+content --requests.params.messages.content.is_error --requests.params.messages.+content --requests.params.messages.role --requests.params.+message --requests.params.model --requests.params.metadata.user_id --requests.params.stop_sequences --requests.params.+stop_sequence --requests.params.stream --requests.params.system.text --requests.params.system.type --requests.params.system.cache_control.type --requests.params.+system --requests.params.temperature --requests.params.tool_choice.type --requests.params.tool_choice.disable_parallel_tool_use --requests.params.tool_choice.name --requests.params.tools.name --requests.params.tools.cache_control.type --requests.params.tools.description --requests.params.+tool --requests.params.top_k --requests.params.top_p --+request"
        ;;
      messages.batches.retrieve)
        flags="--message-batch-id"
        ;;
      messages.batches.list)
        flags="--after-id --before-id --limit"
        ;;
      messages.batches.delete)
        flags="--message-batch-id"
        ;;
      messages.batches.cancel)
        flags="--message-batch-id"
        ;;
      models.retrieve)
        flags="--model-id"
        ;;
      models.list)
        flags="--after-id --before-id --limit"
        ;;
      *)
        # Unknown subcommand
        return
        ;;
    esac

    local cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$COMP_CWORD" -eq 2 || $cur == -* ]] ; then
        COMPREPLY=( $(compgen -W "$flags" -- $cur) )
        return 0
    fi

    local prev="${COMP_WORDS[COMP_CWORD-1]}"
    case "$subcommand" in
      completions.create)
        case "$prev" in
          --model)
            COMPREPLY=( $(compgen -W "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0" -- $cur) )
            ;;
        esac
        ;;
      messages.create)
        case "$prev" in
          --messages.content.type)
            COMPREPLY=( $(compgen -W "text image tool_use tool_result document" -- $cur) )
            ;;
          --messages.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp application/pdf" -- $cur) )
            ;;
          --messages.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.content.type)
            COMPREPLY=( $(compgen -W "text image" -- $cur) )
            ;;
          --messages.content.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.role)
            COMPREPLY=( $(compgen -W "user assistant" -- $cur) )
            ;;
          --model)
            COMPREPLY=( $(compgen -W "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0" -- $cur) )
            ;;
          --system.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --system.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --tool-choice.type)
            COMPREPLY=( $(compgen -W "auto any tool" -- $cur) )
            ;;
          --tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
      messages.count_tokens)
        case "$prev" in
          --messages.content.type)
            COMPREPLY=( $(compgen -W "text image tool_use tool_result document" -- $cur) )
            ;;
          --messages.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp application/pdf" -- $cur) )
            ;;
          --messages.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.content.type)
            COMPREPLY=( $(compgen -W "text image" -- $cur) )
            ;;
          --messages.content.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.role)
            COMPREPLY=( $(compgen -W "user assistant" -- $cur) )
            ;;
          --model)
            COMPREPLY=( $(compgen -W "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0" -- $cur) )
            ;;
          --system.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --system.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --tool-choice.type)
            COMPREPLY=( $(compgen -W "auto any tool" -- $cur) )
            ;;
          --tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
      messages.batches.create)
        case "$prev" in
          --requests.params.messages.content.type)
            COMPREPLY=( $(compgen -W "text image tool_use tool_result document" -- $cur) )
            ;;
          --requests.params.messages.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp application/pdf" -- $cur) )
            ;;
          --requests.params.messages.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --requests.params.messages.content.content.type)
            COMPREPLY=( $(compgen -W "text image" -- $cur) )
            ;;
          --requests.params.messages.content.content.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.content.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --requests.params.messages.content.content.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --requests.params.messages.role)
            COMPREPLY=( $(compgen -W "user assistant" -- $cur) )
            ;;
          --requests.params.model)
            COMPREPLY=( $(compgen -W "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0" -- $cur) )
            ;;
          --requests.params.system.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --requests.params.system.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.tool_choice.type)
            COMPREPLY=( $(compgen -W "auto any tool" -- $cur) )
            ;;
          --requests.params.tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
    esac
}
complete -F _main anthropic-cli