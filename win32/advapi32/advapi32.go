// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in advapi32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package advapi32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("advapi32.dll", false, EntryPoints)
	outside.AddEPs("advapi32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"A_SHAFinal",
	"A_SHAInit",
	"A_SHAUpdate",
	"AbortSystemShutdownA",
	"AccessCheck",
	"AccessCheckAndAuditAlarmA",
	"AccessCheckByType",
	"AccessCheckByTypeAndAuditAlarmA",
	"AccessCheckByTypeResultList",
	"AccessCheckByTypeResultListAndAuditAlarmA",
	"AccessCheckByTypeResultListAndAuditAlarmByHandleA",
	"AddAccessAllowedAce",
	"AddAccessAllowedAceEx",
	"AddAccessAllowedObjectAce",
	"AddAccessDeniedAce",
	"AddAccessDeniedAceEx",
	"AddAccessDeniedObjectAce",
	"AddAce",
	"AddAuditAccessAce",
	"AddAuditAccessAceEx",
	"AddAuditAccessObjectAce",
	"AddUsersToEncryptedFile",
	"AdjustTokenGroups",
	"AdjustTokenPrivileges",
	"AllocateAndInitializeSid",
	"AllocateLocallyUniqueId",
	"AreAllAccessesGranted",
	"AreAnyAccessesGranted",
	"BackupEventLogA",
	"BuildExplicitAccessWithNameA",
	"BuildImpersonateExplicitAccessWithNameA",
	"BuildImpersonateTrusteeA",
	"BuildSecurityDescriptorA",
	"BuildTrusteeWithNameA",
	"BuildTrusteeWithObjectsAndNameA",
	"BuildTrusteeWithObjectsAndSidA",
	"BuildTrusteeWithSidA",
	"CancelOverlappedAccess",
	"ChangeServiceConfig2A",
	"ChangeServiceConfigA",
	"CheckTokenMembership",
	"ClearEventLogA",
	"CloseCodeAuthzLevel",
	"CloseEncryptedFileRaw",
	"CloseEventLog",
	"CloseServiceHandle",
	"CloseTrace",
	"CommandLineFromMsiDescriptor",
	"ComputeAccessTokenFromCodeAuthzLevel",
	"ControlService",
	"ControlTraceA",
	"ConvertAccessToSecurityDescriptorA",
	"ConvertSDToStringSDRootDomainA",
	"ConvertSecurityDescriptorToAccessA",
	"ConvertSecurityDescriptorToAccessNamedA",
	"ConvertSecurityDescriptorToStringSecurityDescriptorA",
	"ConvertSidToStringSidA",
	"ConvertStringSDToSDDomainA",
	"ConvertStringSDToSDRootDomainA",
	"ConvertStringSecurityDescriptorToSecurityDescriptorA",
	"ConvertStringSidToSidA",
	"ConvertToAutoInheritPrivateObjectSecurity",
	"CopySid",
	"CreateCodeAuthzLevel",
	"CreatePrivateObjectSecurity",
	"CreatePrivateObjectSecurityEx",
	"CreatePrivateObjectSecurityWithMultipleInheritance",
	"CreateProcessAsUserA",
	"CreateProcessAsUserSecure",
	"CreateRestrictedToken",
	"CreateServiceA",
	"CreateTraceInstanceId",
	"CreateWellKnownSid",
	"CredDeleteA",
	"CredEnumerateA",
	"CredFree",
	"CredGetSessionTypes",
	"CredGetTargetInfoA",
	"CredIsMarshaledCredentialA",
	"CredMarshalCredentialA",
	"CredProfileLoaded",
	"CredReadA",
	"CredReadDomainCredentialsA",
	"CredRenameA",
	"CredUnmarshalCredentialA",
	"CredWriteA",
	"CredWriteDomainCredentialsA",
	"CredpConvertCredential",
	"CredpConvertTargetInfo",
	"CredpDecodeCredential",
	"CredpEncodeCredential",
	"CryptAcquireContextA",
	"CryptContextAddRef",
	"CryptCreateHash",
	"CryptDecrypt",
	"CryptDeriveKey",
	"CryptDestroyHash",
	"CryptDestroyKey",
	"CryptDuplicateHash",
	"CryptDuplicateKey",
	"CryptEncrypt",
	"CryptEnumProviderTypesA",
	"CryptEnumProvidersA",
	"CryptExportKey",
	"CryptGenKey",
	"CryptGenRandom",
	"CryptGetDefaultProviderA",
	"CryptGetHashParam",
	"CryptGetKeyParam",
	"CryptGetProvParam",
	"CryptGetUserKey",
	"CryptHashData",
	"CryptHashSessionKey",
	"CryptImportKey",
	"CryptReleaseContext",
	"CryptSetHashParam",
	"CryptSetKeyParam",
	"CryptSetProvParam",
	"CryptSetProviderA",
	"CryptSetProviderExA",
	"CryptSignHashA",
	"CryptVerifySignatureA",
	"DecryptFileA",
	"DeleteAce",
	"DeleteService",
	"DeregisterEventSource",
	"DestroyPrivateObjectSecurity",
	"DuplicateEncryptionInfoFile",
	"DuplicateToken",
	"DuplicateTokenEx",
	"ElfBackupEventLogFileA",
	"ElfChangeNotify",
	"ElfClearEventLogFileA",
	"ElfCloseEventLog",
	"ElfDeregisterEventSource",
	"ElfFlushEventLog",
	"ElfNumberOfRecords",
	"ElfOldestRecord",
	"ElfOpenBackupEventLogA",
	"ElfOpenEventLogA",
	"ElfReadEventLogA",
	"ElfRegisterEventSourceA",
	"ElfReportEventA",
	"EnableTrace",
	"EncryptFileA",
	"EncryptedFileKeyInfo",
	"EncryptionDisable",
	"EnumDependentServicesA",
	"EnumServicesStatusA",
	"EnumServicesStatusExA",
	"EnumerateTraceGuids",
	"EqualDomainSid",
	"EqualPrefixSid",
	"EqualSid",
	"FileEncryptionStatusA",
	"FindFirstFreeAce",
	"FlushTraceA",
	"FreeEncryptedFileKeyInfo",
	"FreeEncryptionCertificateHashList",
	"FreeInheritedFromArray",
	"FreeSid",
	"GetAccessPermissionsForObjectA",
	"GetAce",
	"GetAclInformation",
	"GetAuditedPermissionsFromAclA",
	"GetCurrentHwProfileA",
	"GetEffectiveRightsFromAclA",
	"GetEventLogInformation",
	"GetExplicitEntriesFromAclA",
	"GetFileSecurityA",
	"GetInheritanceSourceA",
	"GetKernelObjectSecurity",
	"GetLengthSid",
	"GetLocalManagedApplicationData",
	"GetLocalManagedApplications",
	"GetManagedApplicationCategories",
	"GetManagedApplications",
	"GetMultipleTrusteeA",
	"GetMultipleTrusteeOperationA",
	"GetNamedSecurityInfoA",
	"GetNamedSecurityInfoExA",
	"GetNumberOfEventLogRecords",
	"GetOldestEventLogRecord",
	"GetOverlappedAccessResults",
	"GetPrivateObjectSecurity",
	"GetSecurityDescriptorControl",
	"GetSecurityDescriptorDacl",
	"GetSecurityDescriptorGroup",
	"GetSecurityDescriptorLength",
	"GetSecurityDescriptorOwner",
	"GetSecurityDescriptorRMControl",
	"GetSecurityDescriptorSacl",
	"GetSecurityInfo",
	"GetSecurityInfoExA",
	"GetServiceDisplayNameA",
	"GetServiceKeyNameA",
	"GetSidIdentifierAuthority",
	"GetSidLengthRequired",
	"GetSidSubAuthority",
	"GetSidSubAuthorityCount",
	"GetTokenInformation",
	"GetTraceEnableFlags",
	"GetTraceEnableLevel",
	"GetTraceLoggerHandle",
	"GetTrusteeFormA",
	"GetTrusteeNameA",
	"GetTrusteeTypeA",
	"GetUserNameA",
	"GetWindowsAccountDomainSid",
	"I_ScIsSecurityProcess",
	"I_ScPnPGetServiceName",
	"I_ScSendTSMessage",
	"I_ScSetServiceBitsA",
	"ImpersonateAnonymousToken",
	"ImpersonateLoggedOnUser",
	"ImpersonateNamedPipeClient",
	"ImpersonateSelf",
	"InitializeAcl",
	"InitializeSecurityDescriptor",
	"InitializeSid",
	"InitiateSystemShutdownA",
	"InitiateSystemShutdownExA",
	"InstallApplication",
	"IsTextUnicode",
	"IsTokenRestricted",
	"IsTokenUntrusted",
	"IsValidAcl",
	"IsValidSecurityDescriptor",
	"IsValidSid",
	"IsWellKnownSid",
	"LockServiceDatabase",
	"LogonUserA",
	"LogonUserExA",
	"LookupAccountNameA",
	"LookupAccountSidA",
	"LookupPrivilegeDisplayNameA",
	"LookupPrivilegeNameA",
	"LookupPrivilegeValueA",
	"LookupSecurityDescriptorPartsA",
	"LsaAddAccountRights",
	"LsaAddPrivilegesToAccount",
	"LsaClearAuditLog",
	"LsaClose",
	"LsaCreateAccount",
	"LsaCreateSecret",
	"LsaCreateTrustedDomain",
	"LsaCreateTrustedDomainEx",
	"LsaDelete",
	"LsaDeleteTrustedDomain",
	"LsaEnumerateAccountRights",
	"LsaEnumerateAccounts",
	"LsaEnumerateAccountsWithUserRight",
	"LsaEnumeratePrivileges",
	"LsaEnumeratePrivilegesOfAccount",
	"LsaEnumerateTrustedDomains",
	"LsaEnumerateTrustedDomainsEx",
	"LsaFreeMemory",
	"LsaGetQuotasForAccount",
	"LsaGetRemoteUserName",
	"LsaGetSystemAccessAccount",
	"LsaGetUserName",
	"LsaICLookupNames",
	"LsaICLookupNamesWithCreds",
	"LsaICLookupSids",
	"LsaICLookupSidsWithCreds",
	"LsaLookupNames",
	"LsaLookupNames2",
	"LsaLookupPrivilegeDisplayName",
	"LsaLookupPrivilegeName",
	"LsaLookupPrivilegeValue",
	"LsaLookupSids",
	"LsaNtStatusToWinError",
	"LsaOpenAccount",
	"LsaOpenPolicy",
	"LsaOpenPolicySce",
	"LsaOpenSecret",
	"LsaOpenTrustedDomain",
	"LsaOpenTrustedDomainByName",
	"LsaQueryDomainInformationPolicy",
	"LsaQueryForestTrustInformation",
	"LsaQueryInfoTrustedDomain",
	"LsaQueryInformationPolicy",
	"LsaQuerySecret",
	"LsaQuerySecurityObject",
	"LsaQueryTrustedDomainInfo",
	"LsaQueryTrustedDomainInfoByName",
	"LsaRemoveAccountRights",
	"LsaRemovePrivilegesFromAccount",
	"LsaRetrievePrivateData",
	"LsaSetDomainInformationPolicy",
	"LsaSetForestTrustInformation",
	"LsaSetInformationPolicy",
	"LsaSetInformationTrustedDomain",
	"LsaSetQuotasForAccount",
	"LsaSetSecret",
	"LsaSetSecurityObject",
	"LsaSetSystemAccessAccount",
	"LsaSetTrustedDomainInfoByName",
	"LsaSetTrustedDomainInformation",
	"LsaStorePrivateData",
	"MD4Final",
	"MD4Init",
	"MD4Update",
	"MD5Final",
	"MD5Init",
	"MD5Update",
	"MSChapSrvChangePassword",
	"MSChapSrvChangePassword2",
	"MakeAbsoluteSD",
	"MakeAbsoluteSD2",
	"MakeSelfRelativeSD",
	"MapGenericMask",
	"NotifyBootConfigStatus",
	"NotifyChangeEventLog",
	"ObjectCloseAuditAlarmA",
	"ObjectDeleteAuditAlarmA",
	"ObjectOpenAuditAlarmA",
	"ObjectPrivilegeAuditAlarmA",
	"OpenBackupEventLogA",
	"OpenEncryptedFileRawA",
	"OpenEventLogA",
	"OpenProcessToken",
	"OpenSCManagerA",
	"OpenServiceA",
	"OpenThreadToken",
	"OpenTraceA",
	"PrivilegeCheck",
	"PrivilegedServiceAuditAlarmA",
	"ProcessIdleTasks",
	"ProcessTrace",
	"QueryAllTracesA",
	"QueryRecoveryAgentsOnEncryptedFile",
	"QueryServiceConfig2A",
	"QueryServiceConfigA",
	"QueryServiceLockStatusA",
	"QueryServiceObjectSecurity",
	"QueryServiceStatus",
	"QueryServiceStatusEx",
	"QueryTraceA",
	"QueryUsersOnEncryptedFile",
	"QueryWindows31FilesMigration",
	"ReadEncryptedFileRaw",
	"ReadEventLogA",
	"RegCloseKey",
	"RegConnectRegistryA",
	"RegCreateKeyA",
	"RegCreateKeyExA",
	"RegDeleteKeyA",
	"RegDeleteValueA",
	"RegDisablePredefinedCache",
	"RegDisablePredefinedCacheEx",
	"RegEnumKeyA",
	"RegEnumKeyExA",
	"RegEnumValueA",
	"RegFlushKey",
	"RegGetKeySecurity",
	"RegLoadKeyA",
	"RegNotifyChangeKeyValue",
	"RegOpenCurrentUser",
	"RegOpenKeyA",
	"RegOpenKeyExA",
	"RegOpenUserClassesRoot",
	"RegOverridePredefKey",
	"RegQueryInfoKeyA",
	"RegQueryMultipleValuesA",
	"RegQueryValueA",
	"RegQueryValueExA",
	"RegReplaceKeyA",
	"RegRestoreKeyA",
	"RegSaveKeyA",
	"RegSaveKeyExA",
	"RegSetKeySecurity",
	"RegSetValueA",
	"RegSetValueExA",
	"RegUnLoadKeyA",
	"RegisterEventSourceA",
	"RegisterIdleTask",
	"RegisterServiceCtrlHandlerA",
	"RegisterServiceCtrlHandlerExA",
	"RegisterTraceGuidsA",
	"RemoveTraceCallback",
	"RemoveUsersFromEncryptedFile",
	"ReportEventA",
	"RevertToSelf",
	"SaferCloseLevel",
	"SaferComputeTokenFromLevel",
	"SaferCreateLevel",
	"SaferGetLevelInformation",
	"SaferGetPolicyInformation",
	"SaferIdentifyLevel",
	"SaferRecordEventLogEntry",
	"SaferSetLevelInformation",
	"SaferSetPolicyInformation",
	"SaferiChangeRegistryScope",
	"SaferiCompareTokenLevels",
	"SaferiIsExecutableFileType",
	"SaferiPopulateDefaultsInRegistry",
	"SaferiRecordEventLogEntry",
	"SaferiReplaceProcessThreadTokens",
	"SaferiSearchMatchingHashRules",
	"SetAclInformation",
	"SetEntriesInAccessListA",
	"SetEntriesInAclA",
	"SetEntriesInAuditListA",
	"SetFileSecurityA",
	"SetKernelObjectSecurity",
	"SetNamedSecurityInfoA",
	"SetNamedSecurityInfoExA",
	"SetPrivateObjectSecurity",
	"SetPrivateObjectSecurityEx",
	"SetSecurityDescriptorControl",
	"SetSecurityDescriptorDacl",
	"SetSecurityDescriptorGroup",
	"SetSecurityDescriptorOwner",
	"SetSecurityDescriptorRMControl",
	"SetSecurityDescriptorSacl",
	"SetSecurityInfo",
	"SetSecurityInfoExA",
	"SetServiceBits",
	"SetServiceObjectSecurity",
	"SetServiceStatus",
	"SetThreadToken",
	"SetTokenInformation",
	"SetTraceCallback",
	"SetUserFileEncryptionKey",
	"StartServiceA",
	"StartServiceCtrlDispatcherA",
	"StartTraceA",
	"StopTraceA",
	"SynchronizeWindows31FilesAndWindowsNTRegistry",
	"SystemFunction001",
	"SystemFunction002",
	"SystemFunction003",
	"SystemFunction004",
	"SystemFunction005",
	"SystemFunction006",
	"SystemFunction007",
	"SystemFunction008",
	"SystemFunction009",
	"SystemFunction010",
	"SystemFunction011",
	"SystemFunction012",
	"SystemFunction013",
	"SystemFunction014",
	"SystemFunction015",
	"SystemFunction016",
	"SystemFunction017",
	"SystemFunction018",
	"SystemFunction019",
	"SystemFunction020",
	"SystemFunction021",
	"SystemFunction022",
	"SystemFunction023",
	"SystemFunction024",
	"SystemFunction025",
	"SystemFunction026",
	"SystemFunction027",
	"SystemFunction028",
	"SystemFunction029",
	"SystemFunction030",
	"SystemFunction031",
	"SystemFunction032",
	"SystemFunction033",
	"SystemFunction034",
	"SystemFunction035",
	"SystemFunction036",
	"SystemFunction040",
	"SystemFunction041",
	"TraceEvent",
	"TraceEventInstance",
	"TraceMessage",
	"TraceMessageVa",
	"TreeResetNamedSecurityInfoA",
	"TrusteeAccessToObjectA",
	"UninstallApplication",
	"UnlockServiceDatabase",
	"UnregisterIdleTask",
	"UnregisterTraceGuids",
	"UpdateTraceA",
	"WdmWmiServiceMain",
	"WmiCloseBlock",
	"WmiCloseTraceWithCursor",
	"WmiConvertTimestamp",
	"WmiDevInstToInstanceNameA",
	"WmiEnumerateGuids",
	"WmiExecuteMethodA",
	"WmiFileHandleToInstanceNameA",
	"WmiFreeBuffer",
	"WmiGetFirstTraceOffset",
	"WmiGetNextEvent",
	"WmiGetTraceHeader",
	"WmiMofEnumerateResourcesA",
	"WmiNotificationRegistrationA",
	"WmiOpenBlock",
	"WmiOpenTraceWithCursor",
	"WmiParseTraceEvent",
	"WmiQueryAllDataA",
	"WmiQueryAllDataMultipleA",
	"WmiQueryGuidInformation",
	"WmiQuerySingleInstanceA",
	"WmiQuerySingleInstanceMultipleA",
	"WmiReceiveNotificationsA",
	"WmiSetSingleInstanceA",
	"WmiSetSingleItemA",
	"Wow64Win32ApiEntry",
	"WriteEncryptedFileRaw",
}

var UnicodeEntryPoints = outside.EPs{
	"AbortSystemShutdownW",
	"AccessCheckAndAuditAlarmW",
	"AccessCheckByTypeAndAuditAlarmW",
	"AccessCheckByTypeResultListAndAuditAlarmByHandleW",
	"AccessCheckByTypeResultListAndAuditAlarmW",
	"BackupEventLogW",
	"BuildExplicitAccessWithNameW",
	"BuildImpersonateExplicitAccessWithNameW",
	"BuildImpersonateTrusteeW",
	"BuildSecurityDescriptorW",
	"BuildTrusteeWithNameW",
	"BuildTrusteeWithObjectsAndNameW",
	"BuildTrusteeWithObjectsAndSidW",
	"BuildTrusteeWithSidW",
	"ChangeServiceConfig2W",
	"ChangeServiceConfigW",
	"ClearEventLogW",
	"ControlTraceW",
	"ConvertAccessToSecurityDescriptorW",
	"ConvertSDToStringSDRootDomainW",
	"ConvertSecurityDescriptorToAccessNamedW",
	"ConvertSecurityDescriptorToAccessW",
	"ConvertSecurityDescriptorToStringSecurityDescriptorW",
	"ConvertSidToStringSidW",
	"ConvertStringSDToSDDomainW",
	"ConvertStringSDToSDRootDomainW",
	"ConvertStringSecurityDescriptorToSecurityDescriptorW",
	"ConvertStringSidToSidW",
	"CreateProcessAsUserW",
	"CreateProcessWithLogonW",
	"CreateServiceW",
	"CredDeleteW",
	"CredEnumerateW",
	"CredGetTargetInfoW",
	"CredIsMarshaledCredentialW",
	"CredMarshalCredentialW",
	"CredReadDomainCredentialsW",
	"CredReadW",
	"CredRenameW",
	"CredUnmarshalCredentialW",
	"CredWriteDomainCredentialsW",
	"CredWriteW",
	"CryptAcquireContextW",
	"CryptEnumProviderTypesW",
	"CryptEnumProvidersW",
	"CryptGetDefaultProviderW",
	"CryptSetProviderExW",
	"CryptSetProviderW",
	"CryptSignHashW",
	"CryptVerifySignatureW",
	"DecryptFileW",
	"ElfBackupEventLogFileW",
	"ElfClearEventLogFileW",
	"ElfOpenBackupEventLogW",
	"ElfOpenEventLogW",
	"ElfReadEventLogW",
	"ElfRegisterEventSourceW",
	"ElfReportEventW",
	"EncryptFileW",
	"EnumDependentServicesW",
	"EnumServiceGroupW",
	"EnumServicesStatusExW",
	"EnumServicesStatusW",
	"FileEncryptionStatusW",
	"FlushTraceW",
	"GetAccessPermissionsForObjectW",
	"GetAuditedPermissionsFromAclW",
	"GetCurrentHwProfileW",
	"GetEffectiveRightsFromAclW",
	"GetExplicitEntriesFromAclW",
	"GetFileSecurityW",
	"GetInformationCodeAuthzLevelW",
	"GetInformationCodeAuthzPolicyW",
	"GetInheritanceSourceW",
	"GetMultipleTrusteeOperationW",
	"GetMultipleTrusteeW",
	"GetNamedSecurityInfoExW",
	"GetNamedSecurityInfoW",
	"GetSecurityInfoExW",
	"GetServiceDisplayNameW",
	"GetServiceKeyNameW",
	"GetTrusteeFormW",
	"GetTrusteeNameW",
	"GetTrusteeTypeW",
	"GetUserNameW",
	"I_ScGetCurrentGroupStateW",
	"I_ScSetServiceBitsW",
	"IdentifyCodeAuthzLevelW",
	"InitiateSystemShutdownExW",
	"InitiateSystemShutdownW",
	"LogonUserExW",
	"LogonUserW",
	"LookupAccountNameW",
	"LookupAccountSidW",
	"LookupPrivilegeDisplayNameW",
	"LookupPrivilegeNameW",
	"LookupPrivilegeValueW",
	"LookupSecurityDescriptorPartsW",
	"ObjectCloseAuditAlarmW",
	"ObjectDeleteAuditAlarmW",
	"ObjectOpenAuditAlarmW",
	"ObjectPrivilegeAuditAlarmW",
	"OpenBackupEventLogW",
	"OpenEncryptedFileRawW",
	"OpenEventLogW",
	"OpenSCManagerW",
	"OpenServiceW",
	"OpenTraceW",
	"PrivilegedServiceAuditAlarmW",
	"QueryAllTracesW",
	"QueryServiceConfig2W",
	"QueryServiceConfigW",
	"QueryServiceLockStatusW",
	"QueryTraceW",
	"ReadEventLogW",
	"RegConnectRegistryW",
	"RegCreateKeyExW",
	"RegCreateKeyW",
	"RegDeleteKeyW",
	"RegDeleteValueW",
	"RegEnumKeyExW",
	"RegEnumKeyW",
	"RegEnumValueW",
	"RegLoadKeyW",
	"RegOpenKeyExW",
	"RegOpenKeyW",
	"RegQueryInfoKeyW",
	"RegQueryMultipleValuesW",
	"RegQueryValueExW",
	"RegQueryValueW",
	"RegReplaceKeyW",
	"RegRestoreKeyW",
	"RegSaveKeyExW",
	"RegSaveKeyW",
	"RegSetValueExW",
	"RegSetValueW",
	"RegUnLoadKeyW",
	"RegisterEventSourceW",
	"RegisterServiceCtrlHandlerExW",
	"RegisterServiceCtrlHandlerW",
	"RegisterTraceGuidsW",
	"ReportEventW",
	"SetEntriesInAccessListW",
	"SetEntriesInAclW",
	"SetEntriesInAuditListW",
	"SetFileSecurityW",
	"SetInformationCodeAuthzLevelW",
	"SetInformationCodeAuthzPolicyW",
	"SetNamedSecurityInfoExW",
	"SetNamedSecurityInfoW",
	"SetSecurityInfoExW",
	"StartServiceCtrlDispatcherW",
	"StartServiceW",
	"StartTraceW",
	"StopTraceW",
	"TreeResetNamedSecurityInfoW",
	"TrusteeAccessToObjectW",
	"UpdateTraceW",
	"WmiDevInstToInstanceNameW",
	"WmiExecuteMethodW",
	"WmiFileHandleToInstanceNameW",
	"WmiMofEnumerateResourcesW",
	"WmiNotificationRegistrationW",
	"WmiQueryAllDataMultipleW",
	"WmiQueryAllDataW",
	"WmiQuerySingleInstanceMultipleW",
	"WmiQuerySingleInstanceW",
	"WmiReceiveNotificationsW",
	"WmiSetSingleInstanceW",
	"WmiSetSingleItemW",
}
