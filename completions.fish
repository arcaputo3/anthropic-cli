set -l subcommands completions.create messages.create messages.count_tokens messages.batches.create messages.batches.retrieve messages.batches.delete messages.batches.cancel models.retrieve
complete -c anthropic-cli --no-files \
  -n "not __fish_seen_subcommand_from $subcommands" \
  -a "$subcommands"

complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from completions.create" \
  -a "--max-tokens-to-sample --model --prompt --metadata.user_id --stop-sequences --+stop_sequence --temperature --top-k --top-p"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.create" \
  -a "--max-tokens --messages.content.text_block_param.text --messages.content.text_block_param.type --messages.content.text_block_param.cache_control.type --messages.content.image_block_param.source.data --messages.content.image_block_param.source.media_type --messages.content.image_block_param.source.type --messages.content.image_block_param.type --messages.content.image_block_param.cache_control.type --messages.content.tool_use_block_param.id --messages.content.tool_use_block_param.name --messages.content.tool_use_block_param.type --messages.content.tool_use_block_param.cache_control.type --messages.content.tool_result_block_param.tool_use_id --messages.content.tool_result_block_param.type --messages.content.tool_result_block_param.cache_control.type --messages.content.tool_result_block_param.content.text_block_param.text --messages.content.tool_result_block_param.content.text_block_param.type --messages.content.tool_result_block_param.content.text_block_param.cache_control.type --messages.content.tool_result_block_param.content.image_block_param.source.data --messages.content.tool_result_block_param.content.image_block_param.source.media_type --messages.content.tool_result_block_param.content.image_block_param.source.type --messages.content.tool_result_block_param.content.image_block_param.type --messages.content.tool_result_block_param.content.image_block_param.cache_control.type --messages.content.tool_result_block_param.content.+text_block_param --messages.content.tool_result_block_param.content.+image_block_param --messages.content.tool_result_block_param.is_error --messages.content.document_block_param.source.data --messages.content.document_block_param.source.media_type --messages.content.document_block_param.source.type --messages.content.document_block_param.type --messages.content.document_block_param.cache_control.type --messages.content.+text_block_param --messages.content.+image_block_param --messages.content.+tool_use_block_param --messages.content.+tool_result_block_param --messages.content.+document_block_param --messages.role --+message --model --metadata.user_id --stop-sequences --+stop_sequence --system.text --system.type --system.cache_control.type --+system --temperature --tool-choice.tool_choice_auto.type --tool-choice.tool_choice_auto.disable_parallel_tool_use --tool-choice.tool_choice_any.type --tool-choice.tool_choice_any.disable_parallel_tool_use --tool-choice.tool_choice_tool.name --tool-choice.tool_choice_tool.type --tool-choice.tool_choice_tool.disable_parallel_tool_use --tools.name --tools.cache_control.type --tools.description --+tool --top-k --top-p"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.count_tokens" \
  -a "--messages.content.text_block_param.text --messages.content.text_block_param.type --messages.content.text_block_param.cache_control.type --messages.content.image_block_param.source.data --messages.content.image_block_param.source.media_type --messages.content.image_block_param.source.type --messages.content.image_block_param.type --messages.content.image_block_param.cache_control.type --messages.content.tool_use_block_param.id --messages.content.tool_use_block_param.name --messages.content.tool_use_block_param.type --messages.content.tool_use_block_param.cache_control.type --messages.content.tool_result_block_param.tool_use_id --messages.content.tool_result_block_param.type --messages.content.tool_result_block_param.cache_control.type --messages.content.tool_result_block_param.content.text_block_param.text --messages.content.tool_result_block_param.content.text_block_param.type --messages.content.tool_result_block_param.content.text_block_param.cache_control.type --messages.content.tool_result_block_param.content.image_block_param.source.data --messages.content.tool_result_block_param.content.image_block_param.source.media_type --messages.content.tool_result_block_param.content.image_block_param.source.type --messages.content.tool_result_block_param.content.image_block_param.type --messages.content.tool_result_block_param.content.image_block_param.cache_control.type --messages.content.tool_result_block_param.content.+text_block_param --messages.content.tool_result_block_param.content.+image_block_param --messages.content.tool_result_block_param.is_error --messages.content.document_block_param.source.data --messages.content.document_block_param.source.media_type --messages.content.document_block_param.source.type --messages.content.document_block_param.type --messages.content.document_block_param.cache_control.type --messages.content.+text_block_param --messages.content.+image_block_param --messages.content.+tool_use_block_param --messages.content.+tool_result_block_param --messages.content.+document_block_param --messages.role --+message --model --system.union_member_0 --system.union_member_1.text --system.union_member_1.type --system.union_member_1.cache_control.type --system.+union_member_1 --tool-choice.tool_choice_auto.type --tool-choice.tool_choice_auto.disable_parallel_tool_use --tool-choice.tool_choice_any.type --tool-choice.tool_choice_any.disable_parallel_tool_use --tool-choice.tool_choice_tool.name --tool-choice.tool_choice_tool.type --tool-choice.tool_choice_tool.disable_parallel_tool_use --tools.name --tools.cache_control.type --tools.description --+tool"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.batches.create" \
  -a "--requests.custom_id --requests.params.max_tokens --requests.params.messages.content.text_block_param.text --requests.params.messages.content.text_block_param.type --requests.params.messages.content.text_block_param.cache_control.type --requests.params.messages.content.image_block_param.source.data --requests.params.messages.content.image_block_param.source.media_type --requests.params.messages.content.image_block_param.source.type --requests.params.messages.content.image_block_param.type --requests.params.messages.content.image_block_param.cache_control.type --requests.params.messages.content.tool_use_block_param.id --requests.params.messages.content.tool_use_block_param.name --requests.params.messages.content.tool_use_block_param.type --requests.params.messages.content.tool_use_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.tool_use_id --requests.params.messages.content.tool_result_block_param.type --requests.params.messages.content.tool_result_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.text_block_param.text --requests.params.messages.content.tool_result_block_param.content.text_block_param.type --requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.data --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type --requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.type --requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type --requests.params.messages.content.tool_result_block_param.content.+text_block_param --requests.params.messages.content.tool_result_block_param.content.+image_block_param --requests.params.messages.content.tool_result_block_param.is_error --requests.params.messages.content.document_block_param.source.data --requests.params.messages.content.document_block_param.source.media_type --requests.params.messages.content.document_block_param.source.type --requests.params.messages.content.document_block_param.type --requests.params.messages.content.document_block_param.cache_control.type --requests.params.messages.content.+text_block_param --requests.params.messages.content.+image_block_param --requests.params.messages.content.+tool_use_block_param --requests.params.messages.content.+tool_result_block_param --requests.params.messages.content.+document_block_param --requests.params.messages.role --requests.params.+message --requests.params.model --requests.params.metadata.user_id --requests.params.stop_sequences --requests.params.+stop_sequence --requests.params.stream --requests.params.system.text --requests.params.system.type --requests.params.system.cache_control.type --requests.params.+system --requests.params.temperature --requests.params.tool_choice.tool_choice_auto.type --requests.params.tool_choice.tool_choice_auto.disable_parallel_tool_use --requests.params.tool_choice.tool_choice_any.type --requests.params.tool_choice.tool_choice_any.disable_parallel_tool_use --requests.params.tool_choice.tool_choice_tool.name --requests.params.tool_choice.tool_choice_tool.type --requests.params.tool_choice.tool_choice_tool.disable_parallel_tool_use --requests.params.tools.name --requests.params.tools.cache_control.type --requests.params.tools.description --requests.params.+tool --requests.params.top_k --requests.params.top_p --+request"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.batches.retrieve" \
  -a "--message-batch-id"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.batches.delete" \
  -a "--message-batch-id"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.batches.cancel" \
  -a "--message-batch-id"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from models.retrieve" \
  -a "--model-id"

 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from completions.create" \
   -l model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_use_block_param.type \
   -ra "tool_use"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_use_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.type \
   -ra "tool_result"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.document_block_param.source.media_type \
   -ra "application/pdf"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.document_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.document_block_param.type \
   -ra "document"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.document_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.role \
   -ra "user assistant"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l system.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l system.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_auto.type \
   -ra "auto"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_any.type \
   -ra "any"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_tool.type \
   -ra "tool"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tools.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_use_block_param.type \
   -ra "tool_use"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_use_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.type \
   -ra "tool_result"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.tool_result_block_param.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.document_block_param.source.media_type \
   -ra "application/pdf"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.document_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.document_block_param.type \
   -ra "document"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.content.document_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l messages.role \
   -ra "user assistant"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l system.union_member_1.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l system.union_member_1.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l tool-choice.tool_choice_auto.type \
   -ra "auto"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l tool-choice.tool_choice_any.type \
   -ra "any"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l tool-choice.tool_choice_tool.type \
   -ra "tool"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.count_tokens" \
   -l tools.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_use_block_param.type \
   -ra "tool_use"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_use_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.type \
   -ra "tool_result"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.document_block_param.source.media_type \
   -ra "application/pdf"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.document_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.document_block_param.type \
   -ra "document"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.content.document_block_param.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.messages.role \
   -ra "user assistant"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.system.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.system.cache_control.type \
   -ra "ephemeral"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.tool_choice.tool_choice_auto.type \
   -ra "auto"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.tool_choice.tool_choice_any.type \
   -ra "any"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.tool_choice.tool_choice_tool.type \
   -ra "tool"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.batches.create" \
   -l requests.params.tools.cache_control.type \
   -ra "ephemeral"