// Generated from TSqlParser.g4 by ANTLR 4.7.

package tsql // TSqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TSqlParser.
type TSqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TSqlParser#tsql_file.
	VisitTsql_file(ctx *Tsql_fileContext) interface{}

	// Visit a parse tree produced by TSqlParser#batch.
	VisitBatch(ctx *BatchContext) interface{}

	// Visit a parse tree produced by TSqlParser#sql_clauses.
	VisitSql_clauses(ctx *Sql_clausesContext) interface{}

	// Visit a parse tree produced by TSqlParser#sql_clause.
	VisitSql_clause(ctx *Sql_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#dml_clause.
	VisitDml_clause(ctx *Dml_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#ddl_clause.
	VisitDdl_clause(ctx *Ddl_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_statement.
	VisitBackup_statement(ctx *Backup_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#cfl_statement.
	VisitCfl_statement(ctx *Cfl_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#block_statement.
	VisitBlock_statement(ctx *Block_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#goto_statement.
	VisitGoto_statement(ctx *Goto_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#throw_statement.
	VisitThrow_statement(ctx *Throw_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#throw_error_number.
	VisitThrow_error_number(ctx *Throw_error_numberContext) interface{}

	// Visit a parse tree produced by TSqlParser#throw_message.
	VisitThrow_message(ctx *Throw_messageContext) interface{}

	// Visit a parse tree produced by TSqlParser#throw_state.
	VisitThrow_state(ctx *Throw_stateContext) interface{}

	// Visit a parse tree produced by TSqlParser#try_catch_statement.
	VisitTry_catch_statement(ctx *Try_catch_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#waitfor_statement.
	VisitWaitfor_statement(ctx *Waitfor_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#print_statement.
	VisitPrint_statement(ctx *Print_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#raiseerror_statement.
	VisitRaiseerror_statement(ctx *Raiseerror_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#empty_statement.
	VisitEmpty_statement(ctx *Empty_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#another_statement.
	VisitAnother_statement(ctx *Another_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_application_role.
	VisitAlter_application_role(ctx *Alter_application_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_application_role.
	VisitCreate_application_role(ctx *Create_application_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_aggregate.
	VisitDrop_aggregate(ctx *Drop_aggregateContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_application_role.
	VisitDrop_application_role(ctx *Drop_application_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly.
	VisitAlter_assembly(ctx *Alter_assemblyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_start.
	VisitAlter_assembly_start(ctx *Alter_assembly_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_clause.
	VisitAlter_assembly_clause(ctx *Alter_assembly_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_from_clause.
	VisitAlter_assembly_from_clause(ctx *Alter_assembly_from_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_from_clause_start.
	VisitAlter_assembly_from_clause_start(ctx *Alter_assembly_from_clause_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_drop_clause.
	VisitAlter_assembly_drop_clause(ctx *Alter_assembly_drop_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_drop_multiple_files.
	VisitAlter_assembly_drop_multiple_files(ctx *Alter_assembly_drop_multiple_filesContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_drop.
	VisitAlter_assembly_drop(ctx *Alter_assembly_dropContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_add_clause.
	VisitAlter_assembly_add_clause(ctx *Alter_assembly_add_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_asssembly_add_clause_start.
	VisitAlter_asssembly_add_clause_start(ctx *Alter_asssembly_add_clause_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_client_file_clause.
	VisitAlter_assembly_client_file_clause(ctx *Alter_assembly_client_file_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_file_name.
	VisitAlter_assembly_file_name(ctx *Alter_assembly_file_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_file_bits.
	VisitAlter_assembly_file_bits(ctx *Alter_assembly_file_bitsContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_as.
	VisitAlter_assembly_as(ctx *Alter_assembly_asContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_with_clause.
	VisitAlter_assembly_with_clause(ctx *Alter_assembly_with_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_assembly_with.
	VisitAlter_assembly_with(ctx *Alter_assembly_withContext) interface{}

	// Visit a parse tree produced by TSqlParser#client_assembly_specifier.
	VisitClient_assembly_specifier(ctx *Client_assembly_specifierContext) interface{}

	// Visit a parse tree produced by TSqlParser#assembly_option.
	VisitAssembly_option(ctx *Assembly_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#network_file_share.
	VisitNetwork_file_share(ctx *Network_file_shareContext) interface{}

	// Visit a parse tree produced by TSqlParser#network_computer.
	VisitNetwork_computer(ctx *Network_computerContext) interface{}

	// Visit a parse tree produced by TSqlParser#network_file_start.
	VisitNetwork_file_start(ctx *Network_file_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#file_path.
	VisitFile_path(ctx *File_pathContext) interface{}

	// Visit a parse tree produced by TSqlParser#file_directory_path_separator.
	VisitFile_directory_path_separator(ctx *File_directory_path_separatorContext) interface{}

	// Visit a parse tree produced by TSqlParser#local_file.
	VisitLocal_file(ctx *Local_fileContext) interface{}

	// Visit a parse tree produced by TSqlParser#local_drive.
	VisitLocal_drive(ctx *Local_driveContext) interface{}

	// Visit a parse tree produced by TSqlParser#multiple_local_files.
	VisitMultiple_local_files(ctx *Multiple_local_filesContext) interface{}

	// Visit a parse tree produced by TSqlParser#multiple_local_file_start.
	VisitMultiple_local_file_start(ctx *Multiple_local_file_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_assembly.
	VisitCreate_assembly(ctx *Create_assemblyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_assembly.
	VisitDrop_assembly(ctx *Drop_assemblyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_asymmetric_key.
	VisitAlter_asymmetric_key(ctx *Alter_asymmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_asymmetric_key_start.
	VisitAlter_asymmetric_key_start(ctx *Alter_asymmetric_key_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#asymmetric_key_option.
	VisitAsymmetric_key_option(ctx *Asymmetric_key_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#asymmetric_key_option_start.
	VisitAsymmetric_key_option_start(ctx *Asymmetric_key_option_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#asymmetric_key_password_change_option.
	VisitAsymmetric_key_password_change_option(ctx *Asymmetric_key_password_change_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_asymmetric_key.
	VisitCreate_asymmetric_key(ctx *Create_asymmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_asymmetric_key.
	VisitDrop_asymmetric_key(ctx *Drop_asymmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_authorization.
	VisitAlter_authorization(ctx *Alter_authorizationContext) interface{}

	// Visit a parse tree produced by TSqlParser#authorization_grantee.
	VisitAuthorization_grantee(ctx *Authorization_granteeContext) interface{}

	// Visit a parse tree produced by TSqlParser#entity_to.
	VisitEntity_to(ctx *Entity_toContext) interface{}

	// Visit a parse tree produced by TSqlParser#colon_colon.
	VisitColon_colon(ctx *Colon_colonContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_authorization_start.
	VisitAlter_authorization_start(ctx *Alter_authorization_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_authorization_for_sql_database.
	VisitAlter_authorization_for_sql_database(ctx *Alter_authorization_for_sql_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_authorization_for_azure_dw.
	VisitAlter_authorization_for_azure_dw(ctx *Alter_authorization_for_azure_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_authorization_for_parallel_dw.
	VisitAlter_authorization_for_parallel_dw(ctx *Alter_authorization_for_parallel_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#class_type.
	VisitClass_type(ctx *Class_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#class_type_for_sql_database.
	VisitClass_type_for_sql_database(ctx *Class_type_for_sql_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#class_type_for_azure_dw.
	VisitClass_type_for_azure_dw(ctx *Class_type_for_azure_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#class_type_for_parallel_dw.
	VisitClass_type_for_parallel_dw(ctx *Class_type_for_parallel_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_availability_group.
	VisitDrop_availability_group(ctx *Drop_availability_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_availability_group.
	VisitAlter_availability_group(ctx *Alter_availability_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_availability_group_start.
	VisitAlter_availability_group_start(ctx *Alter_availability_group_startContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_availability_group_options.
	VisitAlter_availability_group_options(ctx *Alter_availability_group_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_broker_priority.
	VisitCreate_or_alter_broker_priority(ctx *Create_or_alter_broker_priorityContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_broker_priority.
	VisitDrop_broker_priority(ctx *Drop_broker_priorityContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_certificate.
	VisitAlter_certificate(ctx *Alter_certificateContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_column_encryption_key.
	VisitAlter_column_encryption_key(ctx *Alter_column_encryption_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_column_encryption_key.
	VisitCreate_column_encryption_key(ctx *Create_column_encryption_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_certificate.
	VisitDrop_certificate(ctx *Drop_certificateContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_column_encryption_key.
	VisitDrop_column_encryption_key(ctx *Drop_column_encryption_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_column_master_key.
	VisitDrop_column_master_key(ctx *Drop_column_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_contract.
	VisitDrop_contract(ctx *Drop_contractContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_credential.
	VisitDrop_credential(ctx *Drop_credentialContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_cryptograhic_provider.
	VisitDrop_cryptograhic_provider(ctx *Drop_cryptograhic_providerContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_database.
	VisitDrop_database(ctx *Drop_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_database_audit_specification.
	VisitDrop_database_audit_specification(ctx *Drop_database_audit_specificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_database_scoped_credential.
	VisitDrop_database_scoped_credential(ctx *Drop_database_scoped_credentialContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_default.
	VisitDrop_default(ctx *Drop_defaultContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_endpoint.
	VisitDrop_endpoint(ctx *Drop_endpointContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_external_data_source.
	VisitDrop_external_data_source(ctx *Drop_external_data_sourceContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_external_file_format.
	VisitDrop_external_file_format(ctx *Drop_external_file_formatContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_external_library.
	VisitDrop_external_library(ctx *Drop_external_libraryContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_external_resource_pool.
	VisitDrop_external_resource_pool(ctx *Drop_external_resource_poolContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_external_table.
	VisitDrop_external_table(ctx *Drop_external_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_event_notifications.
	VisitDrop_event_notifications(ctx *Drop_event_notificationsContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_event_session.
	VisitDrop_event_session(ctx *Drop_event_sessionContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_fulltext_catalog.
	VisitDrop_fulltext_catalog(ctx *Drop_fulltext_catalogContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_fulltext_index.
	VisitDrop_fulltext_index(ctx *Drop_fulltext_indexContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_fulltext_stoplist.
	VisitDrop_fulltext_stoplist(ctx *Drop_fulltext_stoplistContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_login.
	VisitDrop_login(ctx *Drop_loginContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_master_key.
	VisitDrop_master_key(ctx *Drop_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_message_type.
	VisitDrop_message_type(ctx *Drop_message_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_partition_function.
	VisitDrop_partition_function(ctx *Drop_partition_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_partition_scheme.
	VisitDrop_partition_scheme(ctx *Drop_partition_schemeContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_queue.
	VisitDrop_queue(ctx *Drop_queueContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_remote_service_binding.
	VisitDrop_remote_service_binding(ctx *Drop_remote_service_bindingContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_resource_pool.
	VisitDrop_resource_pool(ctx *Drop_resource_poolContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_db_role.
	VisitDrop_db_role(ctx *Drop_db_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_route.
	VisitDrop_route(ctx *Drop_routeContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_rule.
	VisitDrop_rule(ctx *Drop_ruleContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_schema.
	VisitDrop_schema(ctx *Drop_schemaContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_search_property_list.
	VisitDrop_search_property_list(ctx *Drop_search_property_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_security_policy.
	VisitDrop_security_policy(ctx *Drop_security_policyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_sequence.
	VisitDrop_sequence(ctx *Drop_sequenceContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_server_audit.
	VisitDrop_server_audit(ctx *Drop_server_auditContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_server_audit_specification.
	VisitDrop_server_audit_specification(ctx *Drop_server_audit_specificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_server_role.
	VisitDrop_server_role(ctx *Drop_server_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_service.
	VisitDrop_service(ctx *Drop_serviceContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_signature.
	VisitDrop_signature(ctx *Drop_signatureContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_statistics_name_azure_dw_and_pdw.
	VisitDrop_statistics_name_azure_dw_and_pdw(ctx *Drop_statistics_name_azure_dw_and_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_symmetric_key.
	VisitDrop_symmetric_key(ctx *Drop_symmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_synonym.
	VisitDrop_synonym(ctx *Drop_synonymContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_user.
	VisitDrop_user(ctx *Drop_userContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_workload_group.
	VisitDrop_workload_group(ctx *Drop_workload_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_xml_schema_collection.
	VisitDrop_xml_schema_collection(ctx *Drop_xml_schema_collectionContext) interface{}

	// Visit a parse tree produced by TSqlParser#disable_trigger.
	VisitDisable_trigger(ctx *Disable_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#enable_trigger.
	VisitEnable_trigger(ctx *Enable_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#truncate_table.
	VisitTruncate_table(ctx *Truncate_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_column_master_key.
	VisitCreate_column_master_key(ctx *Create_column_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_credential.
	VisitAlter_credential(ctx *Alter_credentialContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_credential.
	VisitCreate_credential(ctx *Create_credentialContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_cryptographic_provider.
	VisitAlter_cryptographic_provider(ctx *Alter_cryptographic_providerContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_cryptographic_provider.
	VisitCreate_cryptographic_provider(ctx *Create_cryptographic_providerContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_event_notification.
	VisitCreate_event_notification(ctx *Create_event_notificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_event_session.
	VisitCreate_or_alter_event_session(ctx *Create_or_alter_event_sessionContext) interface{}

	// Visit a parse tree produced by TSqlParser#event_session_predicate_expression.
	VisitEvent_session_predicate_expression(ctx *Event_session_predicate_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#event_session_predicate_factor.
	VisitEvent_session_predicate_factor(ctx *Event_session_predicate_factorContext) interface{}

	// Visit a parse tree produced by TSqlParser#event_session_predicate_leaf.
	VisitEvent_session_predicate_leaf(ctx *Event_session_predicate_leafContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_external_data_source.
	VisitAlter_external_data_source(ctx *Alter_external_data_sourceContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_external_library.
	VisitAlter_external_library(ctx *Alter_external_libraryContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_external_library.
	VisitCreate_external_library(ctx *Create_external_libraryContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_external_resource_pool.
	VisitAlter_external_resource_pool(ctx *Alter_external_resource_poolContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_external_resource_pool.
	VisitCreate_external_resource_pool(ctx *Create_external_resource_poolContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_fulltext_catalog.
	VisitAlter_fulltext_catalog(ctx *Alter_fulltext_catalogContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_fulltext_catalog.
	VisitCreate_fulltext_catalog(ctx *Create_fulltext_catalogContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_fulltext_stoplist.
	VisitAlter_fulltext_stoplist(ctx *Alter_fulltext_stoplistContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_fulltext_stoplist.
	VisitCreate_fulltext_stoplist(ctx *Create_fulltext_stoplistContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_login_sql_server.
	VisitAlter_login_sql_server(ctx *Alter_login_sql_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_login_sql_server.
	VisitCreate_login_sql_server(ctx *Create_login_sql_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_login_azure_sql.
	VisitAlter_login_azure_sql(ctx *Alter_login_azure_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_login_azure_sql.
	VisitCreate_login_azure_sql(ctx *Create_login_azure_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_login_azure_sql_dw_and_pdw.
	VisitAlter_login_azure_sql_dw_and_pdw(ctx *Alter_login_azure_sql_dw_and_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_login_pdw.
	VisitCreate_login_pdw(ctx *Create_login_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_master_key_sql_server.
	VisitAlter_master_key_sql_server(ctx *Alter_master_key_sql_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_master_key_sql_server.
	VisitCreate_master_key_sql_server(ctx *Create_master_key_sql_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_master_key_azure_sql.
	VisitAlter_master_key_azure_sql(ctx *Alter_master_key_azure_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_master_key_azure_sql.
	VisitCreate_master_key_azure_sql(ctx *Create_master_key_azure_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_message_type.
	VisitAlter_message_type(ctx *Alter_message_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_partition_function.
	VisitAlter_partition_function(ctx *Alter_partition_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_partition_scheme.
	VisitAlter_partition_scheme(ctx *Alter_partition_schemeContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_remote_service_binding.
	VisitAlter_remote_service_binding(ctx *Alter_remote_service_bindingContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_remote_service_binding.
	VisitCreate_remote_service_binding(ctx *Create_remote_service_bindingContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_resource_pool.
	VisitCreate_resource_pool(ctx *Create_resource_poolContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_resource_governor.
	VisitAlter_resource_governor(ctx *Alter_resource_governorContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_db_role.
	VisitAlter_db_role(ctx *Alter_db_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_db_role.
	VisitCreate_db_role(ctx *Create_db_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_route.
	VisitCreate_route(ctx *Create_routeContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_rule.
	VisitCreate_rule(ctx *Create_ruleContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_schema_sql.
	VisitAlter_schema_sql(ctx *Alter_schema_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_schema.
	VisitCreate_schema(ctx *Create_schemaContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_schema_azure_sql_dw_and_pdw.
	VisitCreate_schema_azure_sql_dw_and_pdw(ctx *Create_schema_azure_sql_dw_and_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_schema_azure_sql_dw_and_pdw.
	VisitAlter_schema_azure_sql_dw_and_pdw(ctx *Alter_schema_azure_sql_dw_and_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_search_property_list.
	VisitCreate_search_property_list(ctx *Create_search_property_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_security_policy.
	VisitCreate_security_policy(ctx *Create_security_policyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_sequence.
	VisitAlter_sequence(ctx *Alter_sequenceContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_sequence.
	VisitCreate_sequence(ctx *Create_sequenceContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_server_audit.
	VisitAlter_server_audit(ctx *Alter_server_auditContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_server_audit.
	VisitCreate_server_audit(ctx *Create_server_auditContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_server_audit_specification.
	VisitAlter_server_audit_specification(ctx *Alter_server_audit_specificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_server_audit_specification.
	VisitCreate_server_audit_specification(ctx *Create_server_audit_specificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_server_configuration.
	VisitAlter_server_configuration(ctx *Alter_server_configurationContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_server_role.
	VisitAlter_server_role(ctx *Alter_server_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_server_role.
	VisitCreate_server_role(ctx *Create_server_roleContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_server_role_pdw.
	VisitAlter_server_role_pdw(ctx *Alter_server_role_pdwContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_service.
	VisitAlter_service(ctx *Alter_serviceContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_service.
	VisitCreate_service(ctx *Create_serviceContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_service_master_key.
	VisitAlter_service_master_key(ctx *Alter_service_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_symmetric_key.
	VisitAlter_symmetric_key(ctx *Alter_symmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_symmetric_key.
	VisitCreate_symmetric_key(ctx *Create_symmetric_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_synonym.
	VisitCreate_synonym(ctx *Create_synonymContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_user.
	VisitAlter_user(ctx *Alter_userContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_user.
	VisitCreate_user(ctx *Create_userContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_user_azure_sql_dw.
	VisitCreate_user_azure_sql_dw(ctx *Create_user_azure_sql_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_user_azure_sql.
	VisitAlter_user_azure_sql(ctx *Alter_user_azure_sqlContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_workload_group.
	VisitAlter_workload_group(ctx *Alter_workload_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_workload_group.
	VisitCreate_workload_group(ctx *Create_workload_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_xml_schema_collection.
	VisitCreate_xml_schema_collection(ctx *Create_xml_schema_collectionContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_queue.
	VisitCreate_queue(ctx *Create_queueContext) interface{}

	// Visit a parse tree produced by TSqlParser#queue_settings.
	VisitQueue_settings(ctx *Queue_settingsContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_queue.
	VisitAlter_queue(ctx *Alter_queueContext) interface{}

	// Visit a parse tree produced by TSqlParser#queue_action.
	VisitQueue_action(ctx *Queue_actionContext) interface{}

	// Visit a parse tree produced by TSqlParser#queue_rebuild_options.
	VisitQueue_rebuild_options(ctx *Queue_rebuild_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_contract.
	VisitCreate_contract(ctx *Create_contractContext) interface{}

	// Visit a parse tree produced by TSqlParser#conversation_statement.
	VisitConversation_statement(ctx *Conversation_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#message_statement.
	VisitMessage_statement(ctx *Message_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#merge_statement.
	VisitMerge_statement(ctx *Merge_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#merge_matched.
	VisitMerge_matched(ctx *Merge_matchedContext) interface{}

	// Visit a parse tree produced by TSqlParser#merge_not_matched.
	VisitMerge_not_matched(ctx *Merge_not_matchedContext) interface{}

	// Visit a parse tree produced by TSqlParser#delete_statement.
	VisitDelete_statement(ctx *Delete_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#delete_statement_from.
	VisitDelete_statement_from(ctx *Delete_statement_fromContext) interface{}

	// Visit a parse tree produced by TSqlParser#insert_statement.
	VisitInsert_statement(ctx *Insert_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#insert_statement_value.
	VisitInsert_statement_value(ctx *Insert_statement_valueContext) interface{}

	// Visit a parse tree produced by TSqlParser#receive_statement.
	VisitReceive_statement(ctx *Receive_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#time.
	VisitTime(ctx *TimeContext) interface{}

	// Visit a parse tree produced by TSqlParser#update_statement.
	VisitUpdate_statement(ctx *Update_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#output_clause.
	VisitOutput_clause(ctx *Output_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#output_dml_list_elem.
	VisitOutput_dml_list_elem(ctx *Output_dml_list_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#output_column_name.
	VisitOutput_column_name(ctx *Output_column_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_database.
	VisitCreate_database(ctx *Create_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_index.
	VisitCreate_index(ctx *Create_indexContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_procedure.
	VisitCreate_or_alter_procedure(ctx *Create_or_alter_procedureContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_trigger.
	VisitCreate_or_alter_trigger(ctx *Create_or_alter_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_dml_trigger.
	VisitCreate_or_alter_dml_trigger(ctx *Create_or_alter_dml_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#dml_trigger_option.
	VisitDml_trigger_option(ctx *Dml_trigger_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#dml_trigger_operation.
	VisitDml_trigger_operation(ctx *Dml_trigger_operationContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_ddl_trigger.
	VisitCreate_or_alter_ddl_trigger(ctx *Create_or_alter_ddl_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#ddl_trigger_operation.
	VisitDdl_trigger_operation(ctx *Ddl_trigger_operationContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_or_alter_function.
	VisitCreate_or_alter_function(ctx *Create_or_alter_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#func_body_returns_select.
	VisitFunc_body_returns_select(ctx *Func_body_returns_selectContext) interface{}

	// Visit a parse tree produced by TSqlParser#func_body_returns_table.
	VisitFunc_body_returns_table(ctx *Func_body_returns_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#func_body_returns_scalar.
	VisitFunc_body_returns_scalar(ctx *Func_body_returns_scalarContext) interface{}

	// Visit a parse tree produced by TSqlParser#procedure_param.
	VisitProcedure_param(ctx *Procedure_paramContext) interface{}

	// Visit a parse tree produced by TSqlParser#procedure_option.
	VisitProcedure_option(ctx *Procedure_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#function_option.
	VisitFunction_option(ctx *Function_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_statistics.
	VisitCreate_statistics(ctx *Create_statisticsContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_table.
	VisitCreate_table(ctx *Create_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_options.
	VisitTable_options(ctx *Table_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_view.
	VisitCreate_view(ctx *Create_viewContext) interface{}

	// Visit a parse tree produced by TSqlParser#view_attribute.
	VisitView_attribute(ctx *View_attributeContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_table.
	VisitAlter_table(ctx *Alter_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_database.
	VisitAlter_database(ctx *Alter_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#database_optionspec.
	VisitDatabase_optionspec(ctx *Database_optionspecContext) interface{}

	// Visit a parse tree produced by TSqlParser#auto_option.
	VisitAuto_option(ctx *Auto_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#change_tracking_option.
	VisitChange_tracking_option(ctx *Change_tracking_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#change_tracking_option_list.
	VisitChange_tracking_option_list(ctx *Change_tracking_option_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#containment_option.
	VisitContainment_option(ctx *Containment_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#cursor_option.
	VisitCursor_option(ctx *Cursor_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#alter_endpoint.
	VisitAlter_endpoint(ctx *Alter_endpointContext) interface{}

	// Visit a parse tree produced by TSqlParser#database_mirroring_option.
	VisitDatabase_mirroring_option(ctx *Database_mirroring_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#mirroring_set_option.
	VisitMirroring_set_option(ctx *Mirroring_set_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#mirroring_partner.
	VisitMirroring_partner(ctx *Mirroring_partnerContext) interface{}

	// Visit a parse tree produced by TSqlParser#mirroring_witness.
	VisitMirroring_witness(ctx *Mirroring_witnessContext) interface{}

	// Visit a parse tree produced by TSqlParser#witness_partner_equal.
	VisitWitness_partner_equal(ctx *Witness_partner_equalContext) interface{}

	// Visit a parse tree produced by TSqlParser#partner_option.
	VisitPartner_option(ctx *Partner_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#witness_option.
	VisitWitness_option(ctx *Witness_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#witness_server.
	VisitWitness_server(ctx *Witness_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#partner_server.
	VisitPartner_server(ctx *Partner_serverContext) interface{}

	// Visit a parse tree produced by TSqlParser#mirroring_host_port_seperator.
	VisitMirroring_host_port_seperator(ctx *Mirroring_host_port_seperatorContext) interface{}

	// Visit a parse tree produced by TSqlParser#partner_server_tcp_prefix.
	VisitPartner_server_tcp_prefix(ctx *Partner_server_tcp_prefixContext) interface{}

	// Visit a parse tree produced by TSqlParser#port_number.
	VisitPort_number(ctx *Port_numberContext) interface{}

	// Visit a parse tree produced by TSqlParser#host.
	VisitHost(ctx *HostContext) interface{}

	// Visit a parse tree produced by TSqlParser#date_correlation_optimization_option.
	VisitDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#db_encryption_option.
	VisitDb_encryption_option(ctx *Db_encryption_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#db_state_option.
	VisitDb_state_option(ctx *Db_state_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#db_update_option.
	VisitDb_update_option(ctx *Db_update_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#db_user_access_option.
	VisitDb_user_access_option(ctx *Db_user_access_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#delayed_durability_option.
	VisitDelayed_durability_option(ctx *Delayed_durability_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#external_access_option.
	VisitExternal_access_option(ctx *External_access_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#hadr_options.
	VisitHadr_options(ctx *Hadr_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#mixed_page_allocation_option.
	VisitMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#parameterization_option.
	VisitParameterization_option(ctx *Parameterization_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#recovery_option.
	VisitRecovery_option(ctx *Recovery_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#service_broker_option.
	VisitService_broker_option(ctx *Service_broker_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#snapshot_option.
	VisitSnapshot_option(ctx *Snapshot_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#sql_option.
	VisitSql_option(ctx *Sql_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#target_recovery_time_option.
	VisitTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#termination.
	VisitTermination(ctx *TerminationContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_index.
	VisitDrop_index(ctx *Drop_indexContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_relational_or_xml_or_spatial_index.
	VisitDrop_relational_or_xml_or_spatial_index(ctx *Drop_relational_or_xml_or_spatial_indexContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_backward_compatible_index.
	VisitDrop_backward_compatible_index(ctx *Drop_backward_compatible_indexContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_procedure.
	VisitDrop_procedure(ctx *Drop_procedureContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_trigger.
	VisitDrop_trigger(ctx *Drop_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_dml_trigger.
	VisitDrop_dml_trigger(ctx *Drop_dml_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_ddl_trigger.
	VisitDrop_ddl_trigger(ctx *Drop_ddl_triggerContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_function.
	VisitDrop_function(ctx *Drop_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_statistics.
	VisitDrop_statistics(ctx *Drop_statisticsContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_table.
	VisitDrop_table(ctx *Drop_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_view.
	VisitDrop_view(ctx *Drop_viewContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_type.
	VisitCreate_type(ctx *Create_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#drop_type.
	VisitDrop_type(ctx *Drop_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#rowset_function_limited.
	VisitRowset_function_limited(ctx *Rowset_function_limitedContext) interface{}

	// Visit a parse tree produced by TSqlParser#openquery.
	VisitOpenquery(ctx *OpenqueryContext) interface{}

	// Visit a parse tree produced by TSqlParser#opendatasource.
	VisitOpendatasource(ctx *OpendatasourceContext) interface{}

	// Visit a parse tree produced by TSqlParser#declare_statement.
	VisitDeclare_statement(ctx *Declare_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#cursor_statement.
	VisitCursor_statement(ctx *Cursor_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_database.
	VisitBackup_database(ctx *Backup_databaseContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_log.
	VisitBackup_log(ctx *Backup_logContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_certificate.
	VisitBackup_certificate(ctx *Backup_certificateContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_master_key.
	VisitBackup_master_key(ctx *Backup_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#backup_service_master_key.
	VisitBackup_service_master_key(ctx *Backup_service_master_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#execute_statement.
	VisitExecute_statement(ctx *Execute_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#execute_statement_arg.
	VisitExecute_statement_arg(ctx *Execute_statement_argContext) interface{}

	// Visit a parse tree produced by TSqlParser#execute_var_string.
	VisitExecute_var_string(ctx *Execute_var_stringContext) interface{}

	// Visit a parse tree produced by TSqlParser#security_statement.
	VisitSecurity_statement(ctx *Security_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_certificate.
	VisitCreate_certificate(ctx *Create_certificateContext) interface{}

	// Visit a parse tree produced by TSqlParser#existing_keys.
	VisitExisting_keys(ctx *Existing_keysContext) interface{}

	// Visit a parse tree produced by TSqlParser#private_key_options.
	VisitPrivate_key_options(ctx *Private_key_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#generate_new_keys.
	VisitGenerate_new_keys(ctx *Generate_new_keysContext) interface{}

	// Visit a parse tree produced by TSqlParser#date_options.
	VisitDate_options(ctx *Date_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#open_key.
	VisitOpen_key(ctx *Open_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#close_key.
	VisitClose_key(ctx *Close_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_key.
	VisitCreate_key(ctx *Create_keyContext) interface{}

	// Visit a parse tree produced by TSqlParser#key_options.
	VisitKey_options(ctx *Key_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#algorithm.
	VisitAlgorithm(ctx *AlgorithmContext) interface{}

	// Visit a parse tree produced by TSqlParser#encryption_mechanism.
	VisitEncryption_mechanism(ctx *Encryption_mechanismContext) interface{}

	// Visit a parse tree produced by TSqlParser#decryption_mechanism.
	VisitDecryption_mechanism(ctx *Decryption_mechanismContext) interface{}

	// Visit a parse tree produced by TSqlParser#grant_permission.
	VisitGrant_permission(ctx *Grant_permissionContext) interface{}

	// Visit a parse tree produced by TSqlParser#set_statement.
	VisitSet_statement(ctx *Set_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#transaction_statement.
	VisitTransaction_statement(ctx *Transaction_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#go_statement.
	VisitGo_statement(ctx *Go_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#use_statement.
	VisitUse_statement(ctx *Use_statementContext) interface{}

	// Visit a parse tree produced by TSqlParser#dbcc_clause.
	VisitDbcc_clause(ctx *Dbcc_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#dbcc_options.
	VisitDbcc_options(ctx *Dbcc_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#execute_clause.
	VisitExecute_clause(ctx *Execute_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#declare_local.
	VisitDeclare_local(ctx *Declare_localContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_type_definition.
	VisitTable_type_definition(ctx *Table_type_definitionContext) interface{}

	// Visit a parse tree produced by TSqlParser#xml_type_definition.
	VisitXml_type_definition(ctx *Xml_type_definitionContext) interface{}

	// Visit a parse tree produced by TSqlParser#xml_schema_collection.
	VisitXml_schema_collection(ctx *Xml_schema_collectionContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_def_table_constraints.
	VisitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_def_table_constraint.
	VisitColumn_def_table_constraint(ctx *Column_def_table_constraintContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_definition.
	VisitColumn_definition(ctx *Column_definitionContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_constraint.
	VisitColumn_constraint(ctx *Column_constraintContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_constraint.
	VisitTable_constraint(ctx *Table_constraintContext) interface{}

	// Visit a parse tree produced by TSqlParser#on_delete.
	VisitOn_delete(ctx *On_deleteContext) interface{}

	// Visit a parse tree produced by TSqlParser#on_update.
	VisitOn_update(ctx *On_updateContext) interface{}

	// Visit a parse tree produced by TSqlParser#index_options.
	VisitIndex_options(ctx *Index_optionsContext) interface{}

	// Visit a parse tree produced by TSqlParser#index_option.
	VisitIndex_option(ctx *Index_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#declare_cursor.
	VisitDeclare_cursor(ctx *Declare_cursorContext) interface{}

	// Visit a parse tree produced by TSqlParser#declare_set_cursor_common.
	VisitDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) interface{}

	// Visit a parse tree produced by TSqlParser#declare_set_cursor_common_partial.
	VisitDeclare_set_cursor_common_partial(ctx *Declare_set_cursor_common_partialContext) interface{}

	// Visit a parse tree produced by TSqlParser#fetch_cursor.
	VisitFetch_cursor(ctx *Fetch_cursorContext) interface{}

	// Visit a parse tree produced by TSqlParser#set_special.
	VisitSet_special(ctx *Set_specialContext) interface{}

	// Visit a parse tree produced by TSqlParser#constant_LOCAL_ID.
	VisitConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) interface{}

	// Visit a parse tree produced by TSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#primitive_expression.
	VisitPrimitive_expression(ctx *Primitive_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#case_expression.
	VisitCase_expression(ctx *Case_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#unary_operator_expression.
	VisitUnary_operator_expression(ctx *Unary_operator_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#bracket_expression.
	VisitBracket_expression(ctx *Bracket_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#constant_expression.
	VisitConstant_expression(ctx *Constant_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by TSqlParser#with_expression.
	VisitWith_expression(ctx *With_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#common_table_expression.
	VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#update_elem.
	VisitUpdate_elem(ctx *Update_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#search_condition_list.
	VisitSearch_condition_list(ctx *Search_condition_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#search_condition.
	VisitSearch_condition(ctx *Search_conditionContext) interface{}

	// Visit a parse tree produced by TSqlParser#search_condition_and.
	VisitSearch_condition_and(ctx *Search_condition_andContext) interface{}

	// Visit a parse tree produced by TSqlParser#search_condition_not.
	VisitSearch_condition_not(ctx *Search_condition_notContext) interface{}

	// Visit a parse tree produced by TSqlParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by TSqlParser#query_expression.
	VisitQuery_expression(ctx *Query_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#sql_union.
	VisitSql_union(ctx *Sql_unionContext) interface{}

	// Visit a parse tree produced by TSqlParser#query_specification.
	VisitQuery_specification(ctx *Query_specificationContext) interface{}

	// Visit a parse tree produced by TSqlParser#top_clause.
	VisitTop_clause(ctx *Top_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#top_percent.
	VisitTop_percent(ctx *Top_percentContext) interface{}

	// Visit a parse tree produced by TSqlParser#top_count.
	VisitTop_count(ctx *Top_countContext) interface{}

	// Visit a parse tree produced by TSqlParser#order_by_clause.
	VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#for_clause.
	VisitFor_clause(ctx *For_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#xml_common_directives.
	VisitXml_common_directives(ctx *Xml_common_directivesContext) interface{}

	// Visit a parse tree produced by TSqlParser#order_by_expression.
	VisitOrder_by_expression(ctx *Order_by_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#group_by_item.
	VisitGroup_by_item(ctx *Group_by_itemContext) interface{}

	// Visit a parse tree produced by TSqlParser#option_clause.
	VisitOption_clause(ctx *Option_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by TSqlParser#optimize_for_arg.
	VisitOptimize_for_arg(ctx *Optimize_for_argContext) interface{}

	// Visit a parse tree produced by TSqlParser#select_list.
	VisitSelect_list(ctx *Select_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#udt_method_arguments.
	VisitUdt_method_arguments(ctx *Udt_method_argumentsContext) interface{}

	// Visit a parse tree produced by TSqlParser#asterisk.
	VisitAsterisk(ctx *AsteriskContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_elem.
	VisitColumn_elem(ctx *Column_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#udt_elem.
	VisitUdt_elem(ctx *Udt_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#expression_elem.
	VisitExpression_elem(ctx *Expression_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#select_list_elem.
	VisitSelect_list_elem(ctx *Select_list_elemContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_sources.
	VisitTable_sources(ctx *Table_sourcesContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_source.
	VisitTable_source(ctx *Table_sourceContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_source_item_joined.
	VisitTable_source_item_joined(ctx *Table_source_item_joinedContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_source_item.
	VisitTable_source_item(ctx *Table_source_itemContext) interface{}

	// Visit a parse tree produced by TSqlParser#open_xml.
	VisitOpen_xml(ctx *Open_xmlContext) interface{}

	// Visit a parse tree produced by TSqlParser#schema_declaration.
	VisitSchema_declaration(ctx *Schema_declarationContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_declaration.
	VisitColumn_declaration(ctx *Column_declarationContext) interface{}

	// Visit a parse tree produced by TSqlParser#change_table.
	VisitChange_table(ctx *Change_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#join_part.
	VisitJoin_part(ctx *Join_partContext) interface{}

	// Visit a parse tree produced by TSqlParser#pivot_clause.
	VisitPivot_clause(ctx *Pivot_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#unpivot_clause.
	VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#full_column_name_list.
	VisitFull_column_name_list(ctx *Full_column_name_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_name_with_hint.
	VisitTable_name_with_hint(ctx *Table_name_with_hintContext) interface{}

	// Visit a parse tree produced by TSqlParser#rowset_function.
	VisitRowset_function(ctx *Rowset_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#bulk_option.
	VisitBulk_option(ctx *Bulk_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#derived_table.
	VisitDerived_table(ctx *Derived_tableContext) interface{}

	// Visit a parse tree produced by TSqlParser#RANKING_WINDOWED_FUNC.
	VisitRANKING_WINDOWED_FUNC(ctx *RANKING_WINDOWED_FUNCContext) interface{}

	// Visit a parse tree produced by TSqlParser#AGGREGATE_WINDOWED_FUNC.
	VisitAGGREGATE_WINDOWED_FUNC(ctx *AGGREGATE_WINDOWED_FUNCContext) interface{}

	// Visit a parse tree produced by TSqlParser#ANALYTIC_WINDOWED_FUNC.
	VisitANALYTIC_WINDOWED_FUNC(ctx *ANALYTIC_WINDOWED_FUNCContext) interface{}

	// Visit a parse tree produced by TSqlParser#SCALAR_FUNCTION.
	VisitSCALAR_FUNCTION(ctx *SCALAR_FUNCTIONContext) interface{}

	// Visit a parse tree produced by TSqlParser#BINARY_CHECKSUM.
	VisitBINARY_CHECKSUM(ctx *BINARY_CHECKSUMContext) interface{}

	// Visit a parse tree produced by TSqlParser#CAST.
	VisitCAST(ctx *CASTContext) interface{}

	// Visit a parse tree produced by TSqlParser#CONVERT.
	VisitCONVERT(ctx *CONVERTContext) interface{}

	// Visit a parse tree produced by TSqlParser#CHECKSUM.
	VisitCHECKSUM(ctx *CHECKSUMContext) interface{}

	// Visit a parse tree produced by TSqlParser#COALESCE.
	VisitCOALESCE(ctx *COALESCEContext) interface{}

	// Visit a parse tree produced by TSqlParser#CURRENT_TIMESTAMP.
	VisitCURRENT_TIMESTAMP(ctx *CURRENT_TIMESTAMPContext) interface{}

	// Visit a parse tree produced by TSqlParser#CURRENT_USER.
	VisitCURRENT_USER(ctx *CURRENT_USERContext) interface{}

	// Visit a parse tree produced by TSqlParser#DATEADD.
	VisitDATEADD(ctx *DATEADDContext) interface{}

	// Visit a parse tree produced by TSqlParser#DATEDIFF.
	VisitDATEDIFF(ctx *DATEDIFFContext) interface{}

	// Visit a parse tree produced by TSqlParser#DATENAME.
	VisitDATENAME(ctx *DATENAMEContext) interface{}

	// Visit a parse tree produced by TSqlParser#DATEPART.
	VisitDATEPART(ctx *DATEPARTContext) interface{}

	// Visit a parse tree produced by TSqlParser#GETDATE.
	VisitGETDATE(ctx *GETDATEContext) interface{}

	// Visit a parse tree produced by TSqlParser#GETUTCDATE.
	VisitGETUTCDATE(ctx *GETUTCDATEContext) interface{}

	// Visit a parse tree produced by TSqlParser#IDENTITY.
	VisitIDENTITY(ctx *IDENTITYContext) interface{}

	// Visit a parse tree produced by TSqlParser#MIN_ACTIVE_ROWVERSION.
	VisitMIN_ACTIVE_ROWVERSION(ctx *MIN_ACTIVE_ROWVERSIONContext) interface{}

	// Visit a parse tree produced by TSqlParser#NULLIF.
	VisitNULLIF(ctx *NULLIFContext) interface{}

	// Visit a parse tree produced by TSqlParser#STUFF.
	VisitSTUFF(ctx *STUFFContext) interface{}

	// Visit a parse tree produced by TSqlParser#SESSION_USER.
	VisitSESSION_USER(ctx *SESSION_USERContext) interface{}

	// Visit a parse tree produced by TSqlParser#SYSTEM_USER.
	VisitSYSTEM_USER(ctx *SYSTEM_USERContext) interface{}

	// Visit a parse tree produced by TSqlParser#ISNULL.
	VisitISNULL(ctx *ISNULLContext) interface{}

	// Visit a parse tree produced by TSqlParser#XML_DATA_TYPE_FUNC.
	VisitXML_DATA_TYPE_FUNC(ctx *XML_DATA_TYPE_FUNCContext) interface{}

	// Visit a parse tree produced by TSqlParser#xml_data_type_methods.
	VisitXml_data_type_methods(ctx *Xml_data_type_methodsContext) interface{}

	// Visit a parse tree produced by TSqlParser#value_method.
	VisitValue_method(ctx *Value_methodContext) interface{}

	// Visit a parse tree produced by TSqlParser#query_method.
	VisitQuery_method(ctx *Query_methodContext) interface{}

	// Visit a parse tree produced by TSqlParser#exist_method.
	VisitExist_method(ctx *Exist_methodContext) interface{}

	// Visit a parse tree produced by TSqlParser#modify_method.
	VisitModify_method(ctx *Modify_methodContext) interface{}

	// Visit a parse tree produced by TSqlParser#nodes_method.
	VisitNodes_method(ctx *Nodes_methodContext) interface{}

	// Visit a parse tree produced by TSqlParser#switch_section.
	VisitSwitch_section(ctx *Switch_sectionContext) interface{}

	// Visit a parse tree produced by TSqlParser#switch_search_condition_section.
	VisitSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) interface{}

	// Visit a parse tree produced by TSqlParser#as_column_alias.
	VisitAs_column_alias(ctx *As_column_aliasContext) interface{}

	// Visit a parse tree produced by TSqlParser#as_table_alias.
	VisitAs_table_alias(ctx *As_table_aliasContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by TSqlParser#with_table_hints.
	VisitWith_table_hints(ctx *With_table_hintsContext) interface{}

	// Visit a parse tree produced by TSqlParser#insert_with_table_hints.
	VisitInsert_with_table_hints(ctx *Insert_with_table_hintsContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_hint.
	VisitTable_hint(ctx *Table_hintContext) interface{}

	// Visit a parse tree produced by TSqlParser#index_value.
	VisitIndex_value(ctx *Index_valueContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_alias_list.
	VisitColumn_alias_list(ctx *Column_alias_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_value_constructor.
	VisitTable_value_constructor(ctx *Table_value_constructorContext) interface{}

	// Visit a parse tree produced by TSqlParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#ranking_windowed_function.
	VisitRanking_windowed_function(ctx *Ranking_windowed_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#aggregate_windowed_function.
	VisitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#analytic_windowed_function.
	VisitAnalytic_windowed_function(ctx *Analytic_windowed_functionContext) interface{}

	// Visit a parse tree produced by TSqlParser#all_distinct_expression.
	VisitAll_distinct_expression(ctx *All_distinct_expressionContext) interface{}

	// Visit a parse tree produced by TSqlParser#over_clause.
	VisitOver_clause(ctx *Over_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#row_or_range_clause.
	VisitRow_or_range_clause(ctx *Row_or_range_clauseContext) interface{}

	// Visit a parse tree produced by TSqlParser#window_frame_extent.
	VisitWindow_frame_extent(ctx *Window_frame_extentContext) interface{}

	// Visit a parse tree produced by TSqlParser#window_frame_bound.
	VisitWindow_frame_bound(ctx *Window_frame_boundContext) interface{}

	// Visit a parse tree produced by TSqlParser#window_frame_preceding.
	VisitWindow_frame_preceding(ctx *Window_frame_precedingContext) interface{}

	// Visit a parse tree produced by TSqlParser#window_frame_following.
	VisitWindow_frame_following(ctx *Window_frame_followingContext) interface{}

	// Visit a parse tree produced by TSqlParser#create_database_option.
	VisitCreate_database_option(ctx *Create_database_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#database_filestream_option.
	VisitDatabase_filestream_option(ctx *Database_filestream_optionContext) interface{}

	// Visit a parse tree produced by TSqlParser#database_file_spec.
	VisitDatabase_file_spec(ctx *Database_file_specContext) interface{}

	// Visit a parse tree produced by TSqlParser#file_group.
	VisitFile_group(ctx *File_groupContext) interface{}

	// Visit a parse tree produced by TSqlParser#file_spec.
	VisitFile_spec(ctx *File_specContext) interface{}

	// Visit a parse tree produced by TSqlParser#entity_name.
	VisitEntity_name(ctx *Entity_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#entity_name_for_azure_dw.
	VisitEntity_name_for_azure_dw(ctx *Entity_name_for_azure_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#entity_name_for_parallel_dw.
	VisitEntity_name_for_parallel_dw(ctx *Entity_name_for_parallel_dwContext) interface{}

	// Visit a parse tree produced by TSqlParser#full_table_name.
	VisitFull_table_name(ctx *Full_table_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#simple_name.
	VisitSimple_name(ctx *Simple_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#func_proc_name.
	VisitFunc_proc_name(ctx *Func_proc_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#ddl_object.
	VisitDdl_object(ctx *Ddl_objectContext) interface{}

	// Visit a parse tree produced by TSqlParser#full_column_name.
	VisitFull_column_name(ctx *Full_column_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_name_list_with_order.
	VisitColumn_name_list_with_order(ctx *Column_name_list_with_orderContext) interface{}

	// Visit a parse tree produced by TSqlParser#column_name_list.
	VisitColumn_name_list(ctx *Column_name_listContext) interface{}

	// Visit a parse tree produced by TSqlParser#cursor_name.
	VisitCursor_name(ctx *Cursor_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#on_off.
	VisitOn_off(ctx *On_offContext) interface{}

	// Visit a parse tree produced by TSqlParser#clustered.
	VisitClustered(ctx *ClusteredContext) interface{}

	// Visit a parse tree produced by TSqlParser#null_notnull.
	VisitNull_notnull(ctx *Null_notnullContext) interface{}

	// Visit a parse tree produced by TSqlParser#null_or_default.
	VisitNull_or_default(ctx *Null_or_defaultContext) interface{}

	// Visit a parse tree produced by TSqlParser#scalar_function_name.
	VisitScalar_function_name(ctx *Scalar_function_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#begin_conversation_timer.
	VisitBegin_conversation_timer(ctx *Begin_conversation_timerContext) interface{}

	// Visit a parse tree produced by TSqlParser#begin_conversation_dialog.
	VisitBegin_conversation_dialog(ctx *Begin_conversation_dialogContext) interface{}

	// Visit a parse tree produced by TSqlParser#contract_name.
	VisitContract_name(ctx *Contract_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#service_name.
	VisitService_name(ctx *Service_nameContext) interface{}

	// Visit a parse tree produced by TSqlParser#end_conversation.
	VisitEnd_conversation(ctx *End_conversationContext) interface{}

	// Visit a parse tree produced by TSqlParser#waitfor_conversation.
	VisitWaitfor_conversation(ctx *Waitfor_conversationContext) interface{}

	// Visit a parse tree produced by TSqlParser#get_conversation.
	VisitGet_conversation(ctx *Get_conversationContext) interface{}

	// Visit a parse tree produced by TSqlParser#queue_id.
	VisitQueue_id(ctx *Queue_idContext) interface{}

	// Visit a parse tree produced by TSqlParser#send_conversation.
	VisitSend_conversation(ctx *Send_conversationContext) interface{}

	// Visit a parse tree produced by TSqlParser#data_type.
	VisitData_type(ctx *Data_typeContext) interface{}

	// Visit a parse tree produced by TSqlParser#default_value.
	VisitDefault_value(ctx *Default_valueContext) interface{}

	// Visit a parse tree produced by TSqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by TSqlParser#sign.
	VisitSign(ctx *SignContext) interface{}

	// Visit a parse tree produced by TSqlParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by TSqlParser#simple_id.
	VisitSimple_id(ctx *Simple_idContext) interface{}

	// Visit a parse tree produced by TSqlParser#comparison_operator.
	VisitComparison_operator(ctx *Comparison_operatorContext) interface{}

	// Visit a parse tree produced by TSqlParser#assignment_operator.
	VisitAssignment_operator(ctx *Assignment_operatorContext) interface{}

	// Visit a parse tree produced by TSqlParser#file_size.
	VisitFile_size(ctx *File_sizeContext) interface{}
}
