// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v3/blocks": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Returns blocks by specified filters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get blocks",
                "operationId": "api_v3_get_blocks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Block workchain.",
                        "name": "workchain",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Block shard id. Must be sent with *workchain*. Example: ` + "`" + `8000000000000000` + "`" + `.",
                        "name": "shard",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Block block seqno. Must be sent with *workchain* and *shard*.",
                        "name": "seqno",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Masterchain block seqno",
                        "name": "mc_seqno",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query blocks with generation UTC timestamp **after** given timestamp.",
                        "name": "start_utime",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query blocks with generation UTC timestamp **before** given timestamp.",
                        "name": "end_utime",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query blocks with ` + "`" + `lt \u003e= start_lt` + "`" + `.",
                        "name": "start_lt",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query blocks with ` + "`" + `lt \u003c= end_lt` + "`" + `.",
                        "name": "end_lt",
                        "in": "query"
                    },
                    {
                        "maximum": 500,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Limit number of queried rows. Use with *offset* to batch read.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Skip first N rows. Use with *limit* to batch read.",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "Sort results by UTC timestamp.",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.BlocksResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        },
        "/api/v3/masterchainBlockShardState": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Get masterchain block shard state. Same as /api/v2/shards.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get masterchain block shard state",
                "operationId": "api_v3_get_masterchainBlockShardState",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Masterchain block seqno.",
                        "name": "seqno",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        },
        "/api/v3/masterchainInfo": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Get first and last indexed block",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get Masterchain Info",
                "operationId": "api_v3_get_masterchain_info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.MasterchainInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        },
        "/api/v3/transactions": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Get transactions by specified filter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get transactions",
                "operationId": "api_v3_get_transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Block workchain.",
                        "name": "workchain",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Block shard id. Must be sent with *workchain*. Example: ` + "`" + `8000000000000000` + "`" + `.",
                        "name": "shard",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Block block seqno. Must be sent with *workchain* and *shard*.",
                        "name": "seqno",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Masterchain block seqno.",
                        "name": "mc_seqno",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "List of account addresses to get transactions. Can be sent in hex, base64 or base64url form.",
                        "name": "account",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Exclude transactions on specified account addresses.",
                        "name": "exclude_account",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Transaction hash.",
                        "name": "hash",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Transaction lt.",
                        "name": "lt",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query transactions with generation UTC timestamp **after** given timestamp.",
                        "name": "start_utime",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query transactions with generation UTC timestamp **before** given timestamp.",
                        "name": "end_utime",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query transactions with ` + "`" + `lt \u003e= start_lt` + "`" + `.",
                        "name": "start_lt",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "Query transactions with ` + "`" + `lt \u003c= end_lt` + "`" + `.",
                        "name": "end_lt",
                        "in": "query"
                    },
                    {
                        "maximum": 500,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Limit number of queried rows. Use with *offset* to batch read.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Skip first N rows. Use with *limit* to batch read.",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "Sort transactions by lt.",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        },
        "/api/v3/transactionsByMasterchainBlock": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Returns transactions from masterchain block and from all shards.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get transactions by Masterchain block",
                "operationId": "api_v3_get_transactions_by_masterchain_block",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Masterchain block seqno.",
                        "name": "seqno",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 500,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Limit number of queried rows. Use with *offset* to batch read.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Skip first N rows. Use with *limit* to batch read.",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "Sort transactions by lt.",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        },
        "/api/v3/transactionsByMessage": {
            "get": {
                "security": [
                    {
                        "APIKeyHeader": []
                    },
                    {
                        "APIKeyQuery": []
                    }
                ],
                "description": "Get transactions whose inbound/outbound message has the specified hash. \\",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Get transactions by message",
                "operationId": "api_v3_get_transactions_by_message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Message hash. Acceptable in hex, base64 and base64url forms.",
                        "name": "msg_hash",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Hash of message body.",
                        "name": "body_hash",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "in",
                            "out"
                        ],
                        "type": "string",
                        "description": "Direction of message.",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "maximum": 500,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Limit number of queried rows. Use with *offset* to batch read.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Skip first N rows. Use with *limit* to batch read.",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/index.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/index.RequestError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "index.ActionPhase": {
            "type": "object",
            "properties": {
                "action_list_hash": {
                    "type": "string"
                },
                "msgs_created": {
                    "type": "integer"
                },
                "no_funds": {
                    "type": "boolean"
                },
                "result_arg": {
                    "type": "integer"
                },
                "result_code": {
                    "type": "integer"
                },
                "skipped_actions": {
                    "type": "integer"
                },
                "spec_actions": {
                    "type": "integer"
                },
                "status_change": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "tot_actions": {
                    "type": "integer"
                },
                "tot_msg_size": {
                    "$ref": "#/definitions/index.MsgSize"
                },
                "total_action_fees": {
                    "type": "integer"
                },
                "total_fwd_fees": {
                    "type": "integer"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "index.Block": {
            "type": "object",
            "properties": {
                "after_merge": {
                    "type": "boolean"
                },
                "after_split": {
                    "type": "boolean"
                },
                "before_split": {
                    "type": "boolean"
                },
                "created_by": {
                    "type": "string"
                },
                "end_lt": {
                    "type": "string",
                    "example": "0"
                },
                "file_hash": {
                    "type": "string"
                },
                "flags": {
                    "type": "integer"
                },
                "gen_catchain_seqno": {
                    "type": "integer"
                },
                "gen_utime": {
                    "type": "string",
                    "example": "0"
                },
                "global_id": {
                    "type": "integer"
                },
                "key_block": {
                    "type": "boolean"
                },
                "master_ref_seqno": {
                    "type": "integer"
                },
                "masterchain_block_ref": {
                    "$ref": "#/definitions/index.BlockId"
                },
                "min_ref_mc_seqno": {
                    "type": "integer"
                },
                "prev_blocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/index.BlockId"
                    }
                },
                "prev_key_block_seqno": {
                    "type": "integer"
                },
                "rand_seed": {
                    "type": "string"
                },
                "root_hash": {
                    "type": "string"
                },
                "seqno": {
                    "type": "integer"
                },
                "shard": {
                    "type": "string",
                    "example": "0"
                },
                "start_lt": {
                    "type": "string",
                    "example": "0"
                },
                "tx_count": {
                    "type": "integer"
                },
                "validator_list_hash_short": {
                    "type": "integer"
                },
                "version": {
                    "type": "integer"
                },
                "vert_seqno": {
                    "type": "integer"
                },
                "vert_seqno_incr": {
                    "type": "boolean"
                },
                "want_merge": {
                    "type": "boolean"
                },
                "want_split": {
                    "type": "boolean"
                },
                "workchain": {
                    "type": "integer"
                }
            }
        },
        "index.BlockId": {
            "type": "object",
            "properties": {
                "seqno": {
                    "type": "integer"
                },
                "shard": {
                    "type": "string",
                    "example": "0"
                },
                "workchain": {
                    "type": "integer"
                }
            }
        },
        "index.BlocksResponse": {
            "type": "object",
            "properties": {
                "blocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/index.Block"
                    }
                }
            }
        },
        "index.BouncePhase": {
            "type": "object",
            "properties": {
                "fwd_fees": {
                    "type": "integer"
                },
                "msg_fees": {
                    "type": "integer"
                },
                "msg_size": {
                    "$ref": "#/definitions/index.MsgSize"
                },
                "req_fwd_fees": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "index.ComputePhase": {
            "type": "object",
            "properties": {
                "account_activated": {
                    "type": "boolean"
                },
                "exit_arg": {
                    "type": "integer"
                },
                "exit_code": {
                    "type": "integer"
                },
                "gas_credit": {
                    "type": "integer"
                },
                "gas_fees": {
                    "type": "integer"
                },
                "gas_limit": {
                    "type": "integer"
                },
                "gas_used": {
                    "type": "integer"
                },
                "mode": {
                    "type": "integer"
                },
                "msg_state_used": {
                    "type": "boolean"
                },
                "reason": {
                    "type": "string"
                },
                "skipped": {
                    "type": "boolean"
                },
                "success": {
                    "type": "boolean"
                },
                "vm_final_state_hash": {
                    "type": "string"
                },
                "vm_init_state_hash": {
                    "type": "string"
                },
                "vm_steps": {
                    "type": "integer"
                }
            }
        },
        "index.CreditPhase": {
            "type": "object",
            "properties": {
                "credit": {
                    "type": "integer"
                },
                "due_fees_collected": {
                    "type": "integer"
                }
            }
        },
        "index.DecodedContent": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "index.MasterchainInfo": {
            "type": "object",
            "properties": {
                "first": {
                    "$ref": "#/definitions/index.Block"
                },
                "last": {
                    "$ref": "#/definitions/index.Block"
                }
            }
        },
        "index.Message": {
            "type": "object",
            "properties": {
                "bounce": {
                    "type": "boolean"
                },
                "bounced": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "integer"
                },
                "created_lt": {
                    "type": "integer"
                },
                "destination": {
                    "type": "string"
                },
                "fwd_fee": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "ihr_disabled": {
                    "type": "boolean"
                },
                "ihr_fee": {
                    "type": "integer"
                },
                "import_fee": {
                    "type": "integer"
                },
                "init_state": {
                    "$ref": "#/definitions/index.MessageContent"
                },
                "message_content": {
                    "$ref": "#/definitions/index.MessageContent"
                },
                "opcode": {
                    "type": "integer"
                },
                "source": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "index.MessageContent": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "decoded": {
                    "$ref": "#/definitions/index.DecodedContent"
                },
                "hash": {
                    "type": "string"
                }
            }
        },
        "index.MsgSize": {
            "type": "object",
            "properties": {
                "bits": {
                    "type": "integer"
                },
                "cells": {
                    "type": "integer"
                }
            }
        },
        "index.RequestError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "index.SplitInfo": {
            "type": "object",
            "properties": {
                "acc_split_depth": {
                    "type": "integer"
                },
                "cur_shard_pfx_len": {
                    "type": "integer"
                },
                "sibling_addr": {
                    "type": "string"
                },
                "this_addr": {
                    "type": "string"
                }
            }
        },
        "index.StoragePhase": {
            "type": "object",
            "properties": {
                "status_change": {
                    "type": "string"
                },
                "storage_fees_collected": {
                    "type": "integer"
                },
                "storage_fees_due": {
                    "type": "integer"
                }
            }
        },
        "index.Transaction": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "account_state_hash_after": {
                    "type": "string"
                },
                "account_state_hash_before": {
                    "type": "string"
                },
                "block_seqno": {
                    "type": "integer"
                },
                "block_shard": {
                    "type": "integer"
                },
                "block_workchain": {
                    "type": "integer"
                },
                "description": {
                    "$ref": "#/definitions/index.TransactionDescr"
                },
                "end_status": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "in_msg": {
                    "$ref": "#/definitions/index.Message"
                },
                "lt": {
                    "type": "integer"
                },
                "mc_block_seqno": {
                    "type": "integer"
                },
                "now": {
                    "type": "integer"
                },
                "orig_status": {
                    "type": "string"
                },
                "out_msgs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/index.Message"
                    }
                },
                "prev_trans_hash": {
                    "type": "string"
                },
                "prev_trans_lt": {
                    "type": "integer"
                },
                "total_fees": {
                    "type": "integer"
                },
                "trace_id": {
                    "type": "string"
                }
            }
        },
        "index.TransactionDescr": {
            "type": "object",
            "properties": {
                "aborted": {
                    "type": "boolean"
                },
                "action": {
                    "$ref": "#/definitions/index.ActionPhase"
                },
                "bounce": {
                    "$ref": "#/definitions/index.BouncePhase"
                },
                "compute_ph": {
                    "$ref": "#/definitions/index.ComputePhase"
                },
                "credit_first": {
                    "type": "boolean"
                },
                "credit_ph": {
                    "$ref": "#/definitions/index.CreditPhase"
                },
                "destroyed": {
                    "type": "boolean"
                },
                "installed": {
                    "type": "boolean"
                },
                "is_tock": {
                    "type": "boolean"
                },
                "split_info": {
                    "$ref": "#/definitions/index.SplitInfo"
                },
                "storage_ph": {
                    "$ref": "#/definitions/index.StoragePhase"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "index.TransactionsResponse": {
            "type": "object",
            "properties": {
                "address_book": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/index.Transaction"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "APIKeyHeader": {
            "type": "apiKey",
            "name": "X-Api-Key",
            "in": "header"
        },
        "APIKeyQuery": {
            "type": "apiKey",
            "name": "api_key",
            "in": "query"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TON Index (Go)",
	Description:      "TON Index collects data from a full node to PostgreSQL database and provides convenient API to an indexed blockchain.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
