#!/usr/bin/env bash

# This also works for zsh: https://zsh.sourceforge.io/Doc/Release/Completion-System.html#Completion-System
_main()
{
    COMPREPLY=()

    local subcommands="completions.create messages.create messages.count_tokens messages.batches.create messages.batches.retrieve messages.batches.delete messages.batches.cancel models.retrieve"

    if [[ "$COMP_CWORD" -eq 1 ]]
    then
      local cur="${COMP_WORDS[COMP_CWORD]}"
      COMPREPLY=($(compgen -W "$subcommands" -- "$cur"))
      return
    fi

    local subcommand="${COMP_WORDS[1]}"
    local flags
    case "$subcommand" in
      completions.create)
        flags="--max-tokens-to-sample --model --prompt --metadata.user_id --stop-sequences --+stop_sequence --temperature --top-k --top-p"
        ;;
      messages.create)
        flags="--max-tokens --messages.content.text_block_param.text --messages.content.text_block_param.type --messages.content.text_block_param.cache_control.type --messages.content.image_block_param.source.data --messages.content.image_block_param.source.media_type --messages.content.image_block_param.source.type --messages.content.image_block_param.type --messages.content.image_block_param.cache_control.type --messages.content.tool_use_block_param.id --messages.content.tool_use_block_param.name --messages.content.tool_use_block_param.type --messages.content.tool_use_block_param.cache_control.type --messages.content.tool_result_block_param.tool_use_id --messages.content.tool_result_block_param.type --messages.content.tool_result_block_param.cache_control.type --messages.content.tool_result_block_param.content.text_block_param.text --messages.content.tool_result_block_param.content.text_block_param.type --messages.content.tool_result_block_param.content.text_block_param.cache_control.type --messages.content.tool_result_block_param.content.image_block_param.source.data --messages.content.tool_result_block_param.content.image_block_param.source.media_type --messages.content.tool_result_block_param.content.image_block_param.source.type --messages.content.tool_result_block_param.content.image_block_param.type --messages.content.tool_result_block_param.content.image_block_param.cache_control.type --messages.content.tool_result_block_param.content.+text_block_param --messages.content.tool_result_block_param.content.+image_block_param --messages.content.tool_result_block_param.is_error --messages.content.document_block_param.source.data --messages.content.document_block_param.source.media_type --messages.content.document_block_param.source.type --messages.content.document_block_param.type --messages.content.document_block_param.cache_control.type --messages.content.+text_block_param --messages.content.+image_block_param --messages.content.+tool_use_block_param --messages.content.+tool_result_block_param --messages.content.+document_block_param --messages.role --+message --model --metadata.user_id --stop-sequences --+stop_sequence --system.text --system.type --system.cache_control.type --+system --temperature --tool-choice.tool_choice_auto.type --tool-choice.tool_choice_auto.disable_parallel_tool_use --tool-choice.tool_choice_any.type --tool-choice.tool_choice_any.disable_parallel_tool_use --tool-choice.tool_choice_tool.name --tool-choice.tool_choice_tool.type --tool-choice.tool_choice_tool.disable_parallel_tool_use --tools.name --tools.cache_control.type --tools.description --+tool --top-k --top-p"
        ;;
      messages.count_tokens)
        flags="--messages.content.text_block_param.text --messages.content.text_block_param.type --messages.content.text_block_param.cache_control.type --messages.content.image_block_param.source.data --messages.content.image_block_param.source.media_type --messages.content.image_block_param.source.type --messages.content.image_block_param.type --messages.content.image_block_param.cache_control.type --messages.content.tool_use_block_param.id --messages.content.tool_use_block_param.name --messages.content.tool_use_block_param.type --messages.content.tool_use_block_param.cache_control.type --messages.content.tool_result_block_param.tool_use_id --messages.content.tool_result_block_param.type --messages.content.tool_result_block_param.cache_control.type --messages.content.tool_result_block_param.content.text_block_param.text --messages.content.tool_result_block_param.content.text_block_param.type --messages.content.tool_result_block_param.content.text_block_param.cache_control.type --messages.content.tool_result_block_param.content.image_block_param.source.data --messages.content.tool_result_block_param.content.image_block_param.source.media_type --messages.content.tool_result_block_param.content.image_block_param.source.type --messages.content.tool_result_block_param.content.image_block_param.type --messages.content.tool_result_block_param.content.image_block_param.cache_control.type --messages.content.tool_result_block_param.content.+text_block_param --messages.content.tool_result_block_param.content.+image_block_param --messages.content.tool_result_block_param.is_error --messages.content.document_block_param.source.data --messages.content.document_block_param.source.media_type --messages.content.document_block_param.source.type --messages.content.document_block_param.type --messages.content.document_block_param.cache_control.type --messages.content.+text_block_param --messages.content.+image_block_param --messages.content.+tool_use_block_param --messages.content.+tool_result_block_param --messages.content.+document_block_param --messages.role --+message --model --system.union_member_0 --system.union_member_1.text --system.union_member_1.type --system.union_member_1.cache_control.type --system.+union_member_1 --tool-choice.tool_choice_auto.type --tool-choice.tool_choice_auto.disable_parallel_tool_use --tool-choice.tool_choice_any.type --tool-choice.tool_choice_any.disable_parallel_tool_use --tool-choice.tool_choice_tool.name --tool-choice.tool_choice_tool.type --tool-choice.tool_choice_tool.disable_parallel_tool_use --tools.name --tools.cache_control.type --tools.description --+tool"
        ;;
      messages.batches.create)
        flags="--requests.custom_id --requests.params.max_tokens --requests.params.messages.content.text_block_param.text --requests.params.messages.content.text_block_param.type --requests.params.messages.content.text_block_param.cache_control.type --requests.params.messages.content.image_block_param.source.data --requests.params.messages.content.image_block_param.source.media_type --requests.params.messages.content.image_block_param.source.type --requests.params.messages.content.image_block_param.type --requests.params.messages.content.image_block_param.cache_control.type --requests.params.messages.content.tool_use_block_param.id --requests.params.messages.content.tool_use_block_param.name --requests.params.messages.content.tool_use_block_param.type --requests.params.messages.content.tool_use_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.tool_use_id --requests.params.messages.content.tool_result_block_param.type --requests.params.messages.content.tool_result_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.text_block_param.text --requests.params.messages.content.tool_result_block_param.content.text_block_param.type --requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.data --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.+text_block_param --requests.params.messages.content.tool_result_block_param.content.+image_block_param --requests.params.messages.content.tool_result_block_param.is_error --requests.params.messages.content.document_block_param.source.data --requests.params.messages.content.document_block_param.source.media_type --requests.params.messages.content.document_block_param.source.type --requests.params.messages.content.document_block_param.type --requests.params.messages.content.document_block_param.cache_control.type --requests.params.messages.content.+text_block_param --requests.params.messages.content.+image_block_param --requests.params.messages.content.+tool_use_block_param --requests.params.messages.content.+tool_result_block_param --requests.params.messages.content.+document_block_param --requests.params.messages.role --requests.params.+message --requests.params.model --requests.params.metadata.user_id --requests.params.stop_sequences --requests.params.+stop_sequence --requests.params.stream --requests.params.system.text --requests.params.system.type --requests.params.system.cache_control.type --requests.params.+system --requests.params.temperature --requests.params.tool_choice.tool_choice_auto.type --requests.params.tool_choice.tool_choice_auto.disable_parallel_tool_use --requests.params.tool_choice.tool_choice_any.type --requests.params.tool_choice.tool_choice_any.disable_parallel_tool_use --requests.params.tool_choice.tool_choice_tool.name --requests.params.tool_choice.tool_choice_tool.type --requests.params.tool_choice.tool_choice_tool.disable_parallel_tool_use --requests.params.tools.name --requests.params.tools.cache_control.type --requests.params.tools.description --requests.params.+tool --requests.params.top_k --requests.params.top_p --+request"
        ;;
      messages.batches.retrieve)
        flags="--message-batch-id"
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
          --messages.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --messages.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --messages.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_use_block_param.type)
            COMPREPLY=( $(compgen -W "tool_use" -- $cur) )
            ;;
          --messages.content.tool_use_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.type)
            COMPREPLY=( $(compgen -W "tool_result" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.document_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "application/pdf" -- $cur) )
            ;;
          --messages.content.document_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.document_block_param.type)
            COMPREPLY=( $(compgen -W "document" -- $cur) )
            ;;
          --messages.content.document_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
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
          --tool-choice.tool_choice_auto.type)
            COMPREPLY=( $(compgen -W "auto" -- $cur) )
            ;;
          --tool-choice.tool_choice_any.type)
            COMPREPLY=( $(compgen -W "any" -- $cur) )
            ;;
          --tool-choice.tool_choice_tool.type)
            COMPREPLY=( $(compgen -W "tool" -- $cur) )
            ;;
          --tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
      messages.count_tokens)
        case "$prev" in
          --messages.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --messages.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --messages.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_use_block_param.type)
            COMPREPLY=( $(compgen -W "tool_use" -- $cur) )
            ;;
          --messages.content.tool_use_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.type)
            COMPREPLY=( $(compgen -W "tool_result" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --messages.content.tool_result_block_param.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.content.document_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "application/pdf" -- $cur) )
            ;;
          --messages.content.document_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --messages.content.document_block_param.type)
            COMPREPLY=( $(compgen -W "document" -- $cur) )
            ;;
          --messages.content.document_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --messages.role)
            COMPREPLY=( $(compgen -W "user assistant" -- $cur) )
            ;;
          --model)
            COMPREPLY=( $(compgen -W "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0" -- $cur) )
            ;;
          --system.union_member_1.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --system.union_member_1.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --tool-choice.tool_choice_auto.type)
            COMPREPLY=( $(compgen -W "auto" -- $cur) )
            ;;
          --tool-choice.tool_choice_any.type)
            COMPREPLY=( $(compgen -W "any" -- $cur) )
            ;;
          --tool-choice.tool_choice_tool.type)
            COMPREPLY=( $(compgen -W "tool" -- $cur) )
            ;;
          --tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
      messages.batches.create)
        case "$prev" in
          --requests.params.messages.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --requests.params.messages.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --requests.params.messages.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --requests.params.messages.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --requests.params.messages.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.tool_use_block_param.type)
            COMPREPLY=( $(compgen -W "tool_use" -- $cur) )
            ;;
          --requests.params.messages.content.tool_use_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.type)
            COMPREPLY=( $(compgen -W "tool_result" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.text_block_param.type)
            COMPREPLY=( $(compgen -W "text" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "image/jpeg image/png image/gif image/webp" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.image_block_param.type)
            COMPREPLY=( $(compgen -W "image" -- $cur) )
            ;;
          --requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
          --requests.params.messages.content.document_block_param.source.media_type)
            COMPREPLY=( $(compgen -W "application/pdf" -- $cur) )
            ;;
          --requests.params.messages.content.document_block_param.source.type)
            COMPREPLY=( $(compgen -W "base64" -- $cur) )
            ;;
          --requests.params.messages.content.document_block_param.type)
            COMPREPLY=( $(compgen -W "document" -- $cur) )
            ;;
          --requests.params.messages.content.document_block_param.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
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
          --requests.params.tool_choice.tool_choice_auto.type)
            COMPREPLY=( $(compgen -W "auto" -- $cur) )
            ;;
          --requests.params.tool_choice.tool_choice_any.type)
            COMPREPLY=( $(compgen -W "any" -- $cur) )
            ;;
          --requests.params.tool_choice.tool_choice_tool.type)
            COMPREPLY=( $(compgen -W "tool" -- $cur) )
            ;;
          --requests.params.tools.cache_control.type)
            COMPREPLY=( $(compgen -W "ephemeral" -- $cur) )
            ;;
        esac
        ;;
    esac
}
complete -F _main anthropic-cli