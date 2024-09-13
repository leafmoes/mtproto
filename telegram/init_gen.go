// Code generated by generate-tl-files; DO NOT EDIT.

package telegram

import tl "github.com/leafmoes/mtproto/internal/encoding/tl"

func init() {
	tl.RegisterObjects(
		&AccessPointRule{},
		&AccountAcceptAuthorizationParams{},
		&AccountAuthorizationForm{},
		&AccountAuthorizations{},
		&AccountAutoDownloadSettings{},
		&AccountCancelPasswordEmailParams{},
		&AccountChangePhoneParams{},
		&AccountCheckUsernameParams{},
		&AccountConfirmPasswordEmailParams{},
		&AccountConfirmPhoneParams{},
		&AccountContentSettings{},
		&AccountCreateThemeParams{},
		&AccountDaysTtl{},
		&AccountDeleteAccountParams{},
		&AccountDeleteSecureValueParams{},
		&AccountFinishTakeoutSessionParams{},
		&AccountGetAccountTtlParams{},
		&AccountGetAllSecureValuesParams{},
		&AccountGetAuthorizationFormParams{},
		&AccountGetAuthorizationsParams{},
		&AccountGetAutoDownloadSettingsParams{},
		&AccountGetContactSignUpNotificationParams{},
		&AccountGetContentSettingsParams{},
		&AccountGetGlobalPrivacySettingsParams{},
		&AccountGetMultiWallPapersParams{},
		&AccountGetNotifyExceptionsParams{},
		&AccountGetNotifySettingsParams{},
		&AccountGetPasswordParams{},
		&AccountGetPasswordSettingsParams{},
		&AccountGetPrivacyParams{},
		&AccountGetSecureValueParams{},
		&AccountGetThemeParams{},
		&AccountGetThemesParams{},
		&AccountGetTmpPasswordParams{},
		&AccountGetWallPaperParams{},
		&AccountGetWallPapersParams{},
		&AccountGetWebAuthorizationsParams{},
		&AccountInitTakeoutSessionParams{},
		&AccountInstallThemeParams{},
		&AccountInstallWallPaperParams{},
		&AccountPassword{},
		&AccountPasswordInputSettings{},
		&AccountPasswordSettings{},
		&AccountPrivacyRules{},
		&AccountRegisterDeviceParams{},
		&AccountReportPeerParams{},
		&AccountResendPasswordEmailParams{},
		&AccountResetAuthorizationParams{},
		&AccountResetNotifySettingsParams{},
		&AccountResetWallPapersParams{},
		&AccountResetWebAuthorizationParams{},
		&AccountResetWebAuthorizationsParams{},
		&AccountSaveAutoDownloadSettingsParams{},
		&AccountSaveSecureValueParams{},
		&AccountSaveThemeParams{},
		&AccountSaveWallPaperParams{},
		&AccountSendChangePhoneCodeParams{},
		&AccountSendConfirmPhoneCodeParams{},
		&AccountSendVerifyEmailCodeParams{},
		&AccountSendVerifyPhoneCodeParams{},
		&AccountSentEmailCode{},
		&AccountSetAccountTtlParams{},
		&AccountSetContactSignUpNotificationParams{},
		&AccountSetContentSettingsParams{},
		&AccountSetGlobalPrivacySettingsParams{},
		&AccountSetPrivacyParams{},
		&AccountTakeout{},
		&AccountThemesNotModified{},
		&AccountThemesObj{},
		&AccountTmpPassword{},
		&AccountUnregisterDeviceParams{},
		&AccountUpdateDeviceLockedParams{},
		&AccountUpdateNotifySettingsParams{},
		&AccountUpdatePasswordSettingsParams{},
		&AccountUpdateProfileParams{},
		&AccountUpdateStatusParams{},
		&AccountUpdateThemeParams{},
		&AccountUpdateUsernameParams{},
		&AccountUploadThemeParams{},
		&AccountUploadWallPaperParams{},
		&AccountVerifyEmailParams{},
		&AccountVerifyPhoneParams{},
		&AccountWallPapersNotModified{},
		&AccountWallPapersObj{},
		&AccountWebAuthorizations{},
		&AuthAcceptLoginTokenParams{},
		&AuthAuthorizationObj{},
		&AuthAuthorizationSignUpRequired{},
		&AuthBindTempAuthKeyParams{},
		&AuthCancelCodeParams{},
		&AuthCheckPasswordParams{},
		&AuthDropTempAuthKeysParams{},
		&AuthExportAuthorizationParams{},
		&AuthExportLoginTokenParams{},
		&AuthExportedAuthorization{},
		&AuthImportAuthorizationParams{},
		&AuthImportBotAuthorizationParams{},
		&AuthImportLoginTokenParams{},
		&AuthLogOutParams{},
		&AuthLoginTokenMigrateTo{},
		&AuthLoginTokenObj{},
		&AuthLoginTokenSuccess{},
		&AuthPasswordRecovery{},
		&AuthRecoverPasswordParams{},
		&AuthRequestPasswordRecoveryParams{},
		&AuthResendCodeParams{},
		&AuthResetAuthorizationsParams{},
		&AuthSendCodeParams{},
		&AuthSentCode{},
		&AuthSentCodeTypeApp{},
		&AuthSentCodeTypeCall{},
		&AuthSentCodeTypeFlashCall{},
		&AuthSentCodeTypeSms{},
		&AuthSignInParams{},
		&AuthSignUpParams{},
		&Authorization{},
		&AutoDownloadSettings{},
		&BankCardOpenURL{},
		&BotCommand{},
		&BotInfo{},
		&BotInlineMediaResult{},
		&BotInlineMessageMediaAuto{},
		&BotInlineMessageMediaContact{},
		&BotInlineMessageMediaGeo{},
		&BotInlineMessageMediaVenue{},
		&BotInlineMessageText{},
		&BotInlineResultObj{},
		&BotsAnswerWebhookJsonQueryParams{},
		&BotsSendCustomRequestParams{},
		&BotsSetBotCommandsParams{},
		&CdnConfig{},
		&CdnPublicKey{},
		&Channel{},
		&ChannelAdminLogEvent{},
		&ChannelAdminLogEventActionChangeAbout{},
		&ChannelAdminLogEventActionChangeLinkedChat{},
		&ChannelAdminLogEventActionChangeLocation{},
		&ChannelAdminLogEventActionChangePhoto{},
		&ChannelAdminLogEventActionChangeStickerSet{},
		&ChannelAdminLogEventActionChangeTitle{},
		&ChannelAdminLogEventActionChangeUsername{},
		&ChannelAdminLogEventActionDefaultBannedRights{},
		&ChannelAdminLogEventActionDeleteMessage{},
		&ChannelAdminLogEventActionEditMessage{},
		&ChannelAdminLogEventActionParticipantInvite{},
		&ChannelAdminLogEventActionParticipantJoin{},
		&ChannelAdminLogEventActionParticipantLeave{},
		&ChannelAdminLogEventActionParticipantToggleAdmin{},
		&ChannelAdminLogEventActionParticipantToggleBan{},
		&ChannelAdminLogEventActionStopPoll{},
		&ChannelAdminLogEventActionToggleInvites{},
		&ChannelAdminLogEventActionTogglePreHistoryHidden{},
		&ChannelAdminLogEventActionToggleSignatures{},
		&ChannelAdminLogEventActionToggleSlowMode{},
		&ChannelAdminLogEventActionUpdatePinned{},
		&ChannelAdminLogEventsFilter{},
		&ChannelForbidden{},
		&ChannelFull{},
		&ChannelLocationEmpty{},
		&ChannelLocationObj{},
		&ChannelMessagesFilterEmpty{},
		&ChannelMessagesFilterObj{},
		&ChannelParticipantAdmin{},
		&ChannelParticipantBanned{},
		&ChannelParticipantCreator{},
		&ChannelParticipantLeft{},
		&ChannelParticipantObj{},
		&ChannelParticipantSelf{},
		&ChannelParticipantsAdmins{},
		&ChannelParticipantsBanned{},
		&ChannelParticipantsBots{},
		&ChannelParticipantsContacts{},
		&ChannelParticipantsKicked{},
		&ChannelParticipantsMentions{},
		&ChannelParticipantsRecent{},
		&ChannelParticipantsSearch{},
		&ChannelsAdminLogResults{},
		&ChannelsChannelParticipant{},
		&ChannelsChannelParticipantsNotModified{},
		&ChannelsChannelParticipantsObj{},
		&ChannelsCheckUsernameParams{},
		&ChannelsCreateChannelParams{},
		&ChannelsDeleteChannelParams{},
		&ChannelsDeleteHistoryParams{},
		&ChannelsDeleteMessagesParams{},
		&ChannelsDeleteUserHistoryParams{},
		&ChannelsEditAdminParams{},
		&ChannelsEditBannedParams{},
		&ChannelsEditCreatorParams{},
		&ChannelsEditLocationParams{},
		&ChannelsEditPhotoParams{},
		&ChannelsEditTitleParams{},
		&ChannelsExportMessageLinkParams{},
		&ChannelsGetAdminLogParams{},
		&ChannelsGetAdminedPublicChannelsParams{},
		&ChannelsGetChannelsParams{},
		&ChannelsGetFullChannelParams{},
		&ChannelsGetGroupsForDiscussionParams{},
		&ChannelsGetInactiveChannelsParams{},
		&ChannelsGetLeftChannelsParams{},
		&ChannelsGetMessagesParams{},
		&ChannelsGetParticipantParams{},
		&ChannelsGetParticipantsParams{},
		&ChannelsInviteToChannelParams{},
		&ChannelsJoinChannelParams{},
		&ChannelsLeaveChannelParams{},
		&ChannelsReadHistoryParams{},
		&ChannelsReadMessageContentsParams{},
		&ChannelsReportSpamParams{},
		&ChannelsSetDiscussionGroupParams{},
		&ChannelsSetStickersParams{},
		&ChannelsTogglePreHistoryHiddenParams{},
		&ChannelsToggleSignaturesParams{},
		&ChannelsToggleSlowModeParams{},
		&ChannelsUpdateUsernameParams{},
		&ChatAdminRights{},
		&ChatBannedRights{},
		&ChatEmpty{},
		&ChatForbidden{},
		&ChatFullObj{},
		&ChatInviteAlready{},
		&ChatInviteEmpty{},
		&ChatInviteExported{},
		&ChatInviteObj{},
		&ChatInvitePeek{},
		&ChatObj{},
		&ChatOnlines{},
		&ChatParticipantAdmin{},
		&ChatParticipantCreator{},
		&ChatParticipantObj{},
		&ChatParticipantsForbidden{},
		&ChatParticipantsObj{},
		&ChatPhotoEmpty{},
		&ChatPhotoObj{},
		&CodeSettings{},
		&Config{},
		&Contact{},
		&ContactStatus{},
		&ContactsAcceptContactParams{},
		&ContactsAddContactParams{},
		&ContactsBlockFromRepliesParams{},
		&ContactsBlockParams{},
		&ContactsBlockedObj{},
		&ContactsBlockedSlice{},
		&ContactsContactsNotModified{},
		&ContactsContactsObj{},
		&ContactsDeleteByPhonesParams{},
		&ContactsDeleteContactsParams{},
		&ContactsFound{},
		&ContactsGetBlockedParams{},
		&ContactsGetContactIDsParams{},
		&ContactsGetContactsParams{},
		&ContactsGetLocatedParams{},
		&ContactsGetSavedParams{},
		&ContactsGetStatusesParams{},
		&ContactsGetTopPeersParams{},
		&ContactsImportContactsParams{},
		&ContactsImportedContacts{},
		&ContactsResetSavedParams{},
		&ContactsResetTopPeerRatingParams{},
		&ContactsResolveUsernameParams{},
		&ContactsResolvedPeer{},
		&ContactsSearchParams{},
		&ContactsToggleTopPeersParams{},
		&ContactsTopPeersDisabled{},
		&ContactsTopPeersNotModified{},
		&ContactsTopPeersObj{},
		&ContactsUnblockParams{},
		&DataJson{},
		&DcOption{},
		&DialogFilter{},
		&DialogFilterSuggested{},
		&DialogFolder{},
		&DialogObj{},
		&DialogPeerFolder{},
		&DialogPeerObj{},
		&DocumentAttributeAnimated{},
		&DocumentAttributeAudio{},
		&DocumentAttributeFilename{},
		&DocumentAttributeHasStickers{},
		&DocumentAttributeImageSize{},
		&DocumentAttributeSticker{},
		&DocumentAttributeVideo{},
		&DocumentEmpty{},
		&DocumentObj{},
		&DraftMessageEmpty{},
		&DraftMessageObj{},
		&EmojiKeywordDeleted{},
		&EmojiKeywordObj{},
		&EmojiKeywordsDifference{},
		&EmojiLanguage{},
		&EmojiURL{},
		&EncryptedChatDiscarded{},
		&EncryptedChatEmpty{},
		&EncryptedChatObj{},
		&EncryptedChatRequested{},
		&EncryptedChatWaiting{},
		&EncryptedFileEmpty{},
		&EncryptedFileObj{},
		&EncryptedMessageObj{},
		&EncryptedMessageService{},
		&Error{},
		&ExportedMessageLink{},
		&FileHash{},
		&FileLocationToBeDeprecated{},
		&Folder{},
		&FolderPeer{},
		&FoldersDeleteFolderParams{},
		&FoldersEditPeerFoldersParams{},
		&Game{},
		&GeoPointEmpty{},
		&GeoPointObj{},
		&GlobalPrivacySettings{},
		&HelpAcceptTermsOfServiceParams{},
		&HelpAppUpdateObj{},
		&HelpConfigSimple{},
		&HelpCountriesListNotModified{},
		&HelpCountriesListObj{},
		&HelpCountry{},
		&HelpCountryCode{},
		&HelpDeepLinkInfoEmpty{},
		&HelpDeepLinkInfoObj{},
		&HelpDismissSuggestionParams{},
		&HelpEditUserInfoParams{},
		&HelpGetAppChangelogParams{},
		&HelpGetAppConfigParams{},
		&HelpGetAppUpdateParams{},
		&HelpGetCdnConfigParams{},
		&HelpGetConfigParams{},
		&HelpGetCountriesListParams{},
		&HelpGetDeepLinkInfoParams{},
		&HelpGetInviteTextParams{},
		&HelpGetNearestDcParams{},
		&HelpGetPassportConfigParams{},
		&HelpGetPromoDataParams{},
		&HelpGetRecentMeUrlsParams{},
		&HelpGetSupportNameParams{},
		&HelpGetSupportParams{},
		&HelpGetTermsOfServiceUpdateParams{},
		&HelpGetUserInfoParams{},
		&HelpHidePromoDataParams{},
		&HelpInviteText{},
		&HelpNoAppUpdate{},
		&HelpPassportConfigNotModified{},
		&HelpPassportConfigObj{},
		&HelpPromoDataEmpty{},
		&HelpPromoDataObj{},
		&HelpRecentMeUrls{},
		&HelpSaveAppLogParams{},
		&HelpSetBotUpdatesStatusParams{},
		&HelpSupport{},
		&HelpSupportName{},
		&HelpTermsOfService{},
		&HelpTermsOfServiceUpdateEmpty{},
		&HelpTermsOfServiceUpdateObj{},
		&HelpUserInfoEmpty{},
		&HelpUserInfoObj{},
		&HighScore{},
		&ImportedContact{},
		&InlineBotSwitchPm{},
		&InputAppEvent{},
		&InputBotInlineMessageGame{},
		&InputBotInlineMessageID{},
		&InputBotInlineMessageMediaAuto{},
		&InputBotInlineMessageMediaContact{},
		&InputBotInlineMessageMediaGeo{},
		&InputBotInlineMessageMediaVenue{},
		&InputBotInlineMessageText{},
		&InputBotInlineResultDocument{},
		&InputBotInlineResultGame{},
		&InputBotInlineResultObj{},
		&InputBotInlineResultPhoto{},
		&InputChannelEmpty{},
		&InputChannelFromMessage{},
		&InputChannelObj{},
		&InputChatPhotoEmpty{},
		&InputChatPhotoObj{},
		&InputChatUploadedPhoto{},
		&InputCheckPasswordEmpty{},
		&InputCheckPasswordSRPObj{},
		&InputClientProxy{},
		&InputDialogPeerFolder{},
		&InputDialogPeerObj{},
		&InputDocumentEmpty{},
		&InputDocumentFileLocation{},
		&InputDocumentObj{},
		&InputEncryptedChat{},
		&InputEncryptedFileBigUploaded{},
		&InputEncryptedFileEmpty{},
		&InputEncryptedFileLocation{},
		&InputEncryptedFileObj{},
		&InputEncryptedFileUploaded{},
		&InputFileBig{},
		&InputFileLocationObj{},
		&InputFileObj{},
		&InputFolderPeer{},
		&InputGameID{},
		&InputGameShortName{},
		&InputGeoPointEmpty{},
		&InputGeoPointObj{},
		&InputKeyboardButtonURLAuth{},
		&InputMediaContact{},
		&InputMediaDice{},
		&InputMediaDocument{},
		&InputMediaDocumentExternal{},
		&InputMediaEmpty{},
		&InputMediaGame{},
		&InputMediaGeoLive{},
		&InputMediaGeoPoint{},
		&InputMediaInvoice{},
		&InputMediaPhoto{},
		&InputMediaPhotoExternal{},
		&InputMediaPoll{},
		&InputMediaUploadedDocument{},
		&InputMediaUploadedPhoto{},
		&InputMediaVenue{},
		&InputMessageCallbackQuery{},
		&InputMessageEntityMentionName{},
		&InputMessageID{},
		&InputMessagePinned{},
		&InputMessageReplyTo{},
		&InputMessagesFilterChatPhotos{},
		&InputMessagesFilterContacts{},
		&InputMessagesFilterDocument{},
		&InputMessagesFilterEmpty{},
		&InputMessagesFilterGeo{},
		&InputMessagesFilterGif{},
		&InputMessagesFilterMusic{},
		&InputMessagesFilterMyMentions{},
		&InputMessagesFilterPhoneCalls{},
		&InputMessagesFilterPhotoVideo{},
		&InputMessagesFilterPhotos{},
		&InputMessagesFilterPinned{},
		&InputMessagesFilterRoundVideo{},
		&InputMessagesFilterRoundVoice{},
		&InputMessagesFilterURL{},
		&InputMessagesFilterVideo{},
		&InputMessagesFilterVoice{},
		&InputNotifyBroadcasts{},
		&InputNotifyChats{},
		&InputNotifyPeerObj{},
		&InputNotifyUsers{},
		&InputPaymentCredentialsAndroidPay{},
		&InputPaymentCredentialsApplePay{},
		&InputPaymentCredentialsObj{},
		&InputPaymentCredentialsSaved{},
		&InputPeerChannel{},
		&InputPeerChannelFromMessage{},
		&InputPeerChat{},
		&InputPeerEmpty{},
		&InputPeerNotifySettings{},
		&InputPeerPhotoFileLocation{},
		&InputPeerSelf{},
		&InputPeerUser{},
		&InputPeerUserFromMessage{},
		&InputPhoneCall{},
		&InputPhoneContact{},
		&InputPhotoEmpty{},
		&InputPhotoFileLocation{},
		&InputPhotoLegacyFileLocation{},
		&InputPhotoObj{},
		&InputPrivacyValueAllowAll{},
		&InputPrivacyValueAllowChatParticipants{},
		&InputPrivacyValueAllowContacts{},
		&InputPrivacyValueAllowUsers{},
		&InputPrivacyValueDisallowAll{},
		&InputPrivacyValueDisallowChatParticipants{},
		&InputPrivacyValueDisallowContacts{},
		&InputPrivacyValueDisallowUsers{},
		&InputReportReasonChildAbuse{},
		&InputReportReasonCopyright{},
		&InputReportReasonGeoIrrelevant{},
		&InputReportReasonOther{},
		&InputReportReasonPornography{},
		&InputReportReasonSpam{},
		&InputReportReasonViolence{},
		&InputSecureFileLocation{},
		&InputSecureFileObj{},
		&InputSecureFileUploaded{},
		&InputSecureValue{},
		&InputSingleMedia{},
		&InputStickerSetAnimatedEmoji{},
		&InputStickerSetDice{},
		&InputStickerSetEmpty{},
		&InputStickerSetID{},
		&InputStickerSetItem{},
		&InputStickerSetShortName{},
		&InputStickerSetThumb{},
		&InputStickeredMediaDocument{},
		&InputStickeredMediaPhoto{},
		&InputTakeoutFileLocation{},
		&InputThemeObj{},
		&InputThemeSettings{},
		&InputThemeSlug{},
		&InputUserEmpty{},
		&InputUserFromMessage{},
		&InputUserObj{},
		&InputUserSelf{},
		&InputWallPaperNoFile{},
		&InputWallPaperObj{},
		&InputWallPaperSlug{},
		&InputWebDocument{},
		&InputWebFileGeoPointLocation{},
		&InputWebFileLocationObj{},
		&Invoice{},
		&IpPortObj{},
		&IpPortSecret{},
		&JsonArray{},
		&JsonBool{},
		&JsonNull{},
		&JsonNumber{},
		&JsonObject{},
		&JsonObjectValue{},
		&JsonString{},
		&KeyboardButtonBuy{},
		&KeyboardButtonCallback{},
		&KeyboardButtonGame{},
		&KeyboardButtonObj{},
		&KeyboardButtonRequestGeoLocation{},
		&KeyboardButtonRequestPhone{},
		&KeyboardButtonRequestPoll{},
		&KeyboardButtonRow{},
		&KeyboardButtonSwitchInline{},
		&KeyboardButtonURL{},
		&KeyboardButtonURLAuth{},
		&LabeledPrice{},
		&LangPackDifference{},
		&LangPackLanguage{},
		&LangPackStringDeleted{},
		&LangPackStringObj{},
		&LangPackStringPluralized{},
		&LangpackGetDifferenceParams{},
		&LangpackGetLangPackParams{},
		&LangpackGetLanguageParams{},
		&LangpackGetLanguagesParams{},
		&LangpackGetStringsParams{},
		&MaskCoords{},
		&MessageActionBotAllowed{},
		&MessageActionChannelCreate{},
		&MessageActionChannelMigrateFrom{},
		&MessageActionChatAddUser{},
		&MessageActionChatCreate{},
		&MessageActionChatDeletePhoto{},
		&MessageActionChatDeleteUser{},
		&MessageActionChatEditPhoto{},
		&MessageActionChatEditTitle{},
		&MessageActionChatJoinedByLink{},
		&MessageActionChatMigrateTo{},
		&MessageActionContactSignUp{},
		&MessageActionCustomAction{},
		&MessageActionEmpty{},
		&MessageActionGameScore{},
		&MessageActionGeoProximityReached{},
		&MessageActionHistoryClear{},
		&MessageActionPaymentSent{},
		&MessageActionPaymentSentMe{},
		&MessageActionPhoneCall{},
		&MessageActionPinMessage{},
		&MessageActionScreenshotTaken{},
		&MessageActionSecureValuesSent{},
		&MessageActionSecureValuesSentMe{},
		&MessageEmpty{},
		&MessageEntityBankCard{},
		&MessageEntityBlockquote{},
		&MessageEntityBold{},
		&MessageEntityBotCommand{},
		&MessageEntityCashtag{},
		&MessageEntityCode{},
		&MessageEntityEmail{},
		&MessageEntityHashtag{},
		&MessageEntityItalic{},
		&MessageEntityMention{},
		&MessageEntityMentionName{},
		&MessageEntityPhone{},
		&MessageEntityPre{},
		&MessageEntityStrike{},
		&MessageEntityTextURL{},
		&MessageEntityURL{},
		&MessageEntityUnderline{},
		&MessageEntityUnknown{},
		&MessageFwdHeader{},
		&MessageInteractionCounters{},
		&MessageMediaContact{},
		&MessageMediaDice{},
		&MessageMediaDocument{},
		&MessageMediaEmpty{},
		&MessageMediaGame{},
		&MessageMediaGeo{},
		&MessageMediaGeoLive{},
		&MessageMediaInvoice{},
		&MessageMediaPhoto{},
		&MessageMediaPoll{},
		&MessageMediaUnsupported{},
		&MessageMediaVenue{},
		&MessageMediaWebPage{},
		&MessageObj{},
		&MessageRange{},
		&MessageReplies{},
		&MessageReplyHeader{},
		&MessageService{},
		&MessageUserVoteInputOption{},
		&MessageUserVoteMultiple{},
		&MessageUserVoteObj{},
		&MessageViews{},
		&MessagesAcceptEncryptionParams{},
		&MessagesAcceptURLAuthParams{},
		&MessagesAddChatUserParams{},
		&MessagesAffectedHistory{},
		&MessagesAffectedMessages{},
		&MessagesAllStickersNotModified{},
		&MessagesAllStickersObj{},
		&MessagesArchivedStickers{},
		&MessagesBotCallbackAnswer{},
		&MessagesBotResults{},
		&MessagesChannelMessages{},
		&MessagesChatFull{},
		&MessagesChatsObj{},
		&MessagesChatsSlice{},
		&MessagesCheckChatInviteParams{},
		&MessagesClearAllDraftsParams{},
		&MessagesClearRecentStickersParams{},
		&MessagesCreateChatParams{},
		&MessagesDeleteChatUserParams{},
		&MessagesDeleteHistoryParams{},
		&MessagesDeleteMessagesParams{},
		&MessagesDeleteScheduledMessagesParams{},
		&MessagesDhConfigNotModified{},
		&MessagesDhConfigObj{},
		&MessagesDialogsNotModified{},
		&MessagesDialogsObj{},
		&MessagesDialogsSlice{},
		&MessagesDiscardEncryptionParams{},
		&MessagesDiscussionMessage{},
		&MessagesEditChatAboutParams{},
		&MessagesEditChatAdminParams{},
		&MessagesEditChatDefaultBannedRightsParams{},
		&MessagesEditChatPhotoParams{},
		&MessagesEditChatTitleParams{},
		&MessagesEditInlineBotMessageParams{},
		&MessagesEditMessageParams{},
		&MessagesExportChatInviteParams{},
		&MessagesFaveStickerParams{},
		&MessagesFavedStickersNotModified{},
		&MessagesFavedStickersObj{},
		&MessagesFeaturedStickersNotModified{},
		&MessagesFeaturedStickersObj{},
		&MessagesForwardMessagesParams{},
		&MessagesFoundStickerSetsNotModified{},
		&MessagesFoundStickerSetsObj{},
		&MessagesGetAllChatsParams{},
		&MessagesGetAllDraftsParams{},
		&MessagesGetAllStickersParams{},
		&MessagesGetArchivedStickersParams{},
		&MessagesGetAttachedStickersParams{},
		&MessagesGetBotCallbackAnswerParams{},
		&MessagesGetChatsParams{},
		&MessagesGetCommonChatsParams{},
		&MessagesGetDhConfigParams{},
		&MessagesGetDialogFiltersParams{},
		&MessagesGetDialogUnreadMarksParams{},
		&MessagesGetDialogsParams{},
		&MessagesGetDiscussionMessageParams{},
		&MessagesGetDocumentByHashParams{},
		&MessagesGetEmojiKeywordsDifferenceParams{},
		&MessagesGetEmojiKeywordsLanguagesParams{},
		&MessagesGetEmojiKeywordsParams{},
		&MessagesGetEmojiURLParams{},
		&MessagesGetFavedStickersParams{},
		&MessagesGetFeaturedStickersParams{},
		&MessagesGetFullChatParams{},
		&MessagesGetGameHighScoresParams{},
		&MessagesGetHistoryParams{},
		&MessagesGetInlineBotResultsParams{},
		&MessagesGetInlineGameHighScoresParams{},
		&MessagesGetMaskStickersParams{},
		&MessagesGetMessageEditDataParams{},
		&MessagesGetMessagesParams{},
		&MessagesGetMessagesViewsParams{},
		&MessagesGetOldFeaturedStickersParams{},
		&MessagesGetOnlinesParams{},
		&MessagesGetPeerDialogsParams{},
		&MessagesGetPeerSettingsParams{},
		&MessagesGetPinnedDialogsParams{},
		&MessagesGetPollResultsParams{},
		&MessagesGetPollVotesParams{},
		&MessagesGetRecentLocationsParams{},
		&MessagesGetRecentStickersParams{},
		&MessagesGetRepliesParams{},
		&MessagesGetSavedGifsParams{},
		&MessagesGetScheduledHistoryParams{},
		&MessagesGetScheduledMessagesParams{},
		&MessagesGetSearchCountersParams{},
		&MessagesGetSplitRangesParams{},
		&MessagesGetStatsURLParams{},
		&MessagesGetStickerSetParams{},
		&MessagesGetStickersParams{},
		&MessagesGetSuggestedDialogFiltersParams{},
		&MessagesGetUnreadMentionsParams{},
		&MessagesGetWebPageParams{},
		&MessagesGetWebPagePreviewParams{},
		&MessagesHidePeerSettingsBarParams{},
		&MessagesHighScores{},
		&MessagesImportChatInviteParams{},
		&MessagesInactiveChats{},
		&MessagesInstallStickerSetParams{},
		&MessagesMarkDialogUnreadParams{},
		&MessagesMessageEditData{},
		&MessagesMessageViews{},
		&MessagesMessagesNotModified{},
		&MessagesMessagesObj{},
		&MessagesMessagesSlice{},
		&MessagesMigrateChatParams{},
		&MessagesPeerDialogs{},
		&MessagesReadDiscussionParams{},
		&MessagesReadEncryptedHistoryParams{},
		&MessagesReadFeaturedStickersParams{},
		&MessagesReadHistoryParams{},
		&MessagesReadMentionsParams{},
		&MessagesReadMessageContentsParams{},
		&MessagesReceivedMessagesParams{},
		&MessagesReceivedQueueParams{},
		&MessagesRecentStickersNotModified{},
		&MessagesRecentStickersObj{},
		&MessagesReorderPinnedDialogsParams{},
		&MessagesReorderStickerSetsParams{},
		&MessagesReportEncryptedSpamParams{},
		&MessagesReportParams{},
		&MessagesReportSpamParams{},
		&MessagesRequestEncryptionParams{},
		&MessagesRequestURLAuthParams{},
		&MessagesSaveDraftParams{},
		&MessagesSaveGifParams{},
		&MessagesSaveRecentStickerParams{},
		&MessagesSavedGifsNotModified{},
		&MessagesSavedGifsObj{},
		&MessagesSearchCounter{},
		&MessagesSearchGlobalParams{},
		&MessagesSearchParams{},
		&MessagesSearchStickerSetsParams{},
		&MessagesSendEncryptedFileParams{},
		&MessagesSendEncryptedParams{},
		&MessagesSendEncryptedServiceParams{},
		&MessagesSendInlineBotResultParams{},
		&MessagesSendMediaParams{},
		&MessagesSendMessageParams{},
		&MessagesSendMultiMediaParams{},
		&MessagesSendScheduledMessagesParams{},
		&MessagesSendScreenshotNotificationParams{},
		&MessagesSendVoteParams{},
		&MessagesSentEncryptedFile{},
		&MessagesSentEncryptedMessageObj{},
		&MessagesSetBotCallbackAnswerParams{},
		&MessagesSetBotPrecheckoutResultsParams{},
		&MessagesSetBotShippingResultsParams{},
		&MessagesSetEncryptedTypingParams{},
		&MessagesSetGameScoreParams{},
		&MessagesSetInlineBotResultsParams{},
		&MessagesSetInlineGameScoreParams{},
		&MessagesSetTypingParams{},
		&MessagesStartBotParams{},
		&MessagesStickerSet{},
		&MessagesStickerSetInstallResultArchive{},
		&MessagesStickerSetInstallResultSuccess{},
		&MessagesStickersNotModified{},
		&MessagesStickersObj{},
		&MessagesToggleDialogPinParams{},
		&MessagesToggleStickerSetsParams{},
		&MessagesUninstallStickerSetParams{},
		&MessagesUnpinAllMessagesParams{},
		&MessagesUpdateDialogFilterParams{},
		&MessagesUpdateDialogFiltersOrderParams{},
		&MessagesUpdatePinnedMessageParams{},
		&MessagesUploadEncryptedFileParams{},
		&MessagesUploadMediaParams{},
		&MessagesVotesList{},
		&NearestDc{},
		&NotifyBroadcasts{},
		&NotifyChats{},
		&NotifyPeerObj{},
		&NotifyUsers{},
		&Page{},
		&PageBlockAnchor{},
		&PageBlockAudio{},
		&PageBlockAuthorDate{},
		&PageBlockBlockquote{},
		&PageBlockChannel{},
		&PageBlockCollage{},
		&PageBlockCover{},
		&PageBlockDetails{},
		&PageBlockDivider{},
		&PageBlockEmbed{},
		&PageBlockEmbedPost{},
		&PageBlockFooter{},
		&PageBlockHeader{},
		&PageBlockKicker{},
		&PageBlockList{},
		&PageBlockMap{},
		&PageBlockOrderedList{},
		&PageBlockParagraph{},
		&PageBlockPhoto{},
		&PageBlockPreformatted{},
		&PageBlockPullquote{},
		&PageBlockRelatedArticles{},
		&PageBlockSlideshow{},
		&PageBlockSubheader{},
		&PageBlockSubtitle{},
		&PageBlockTable{},
		&PageBlockTitle{},
		&PageBlockUnsupported{},
		&PageBlockVideo{},
		&PageCaption{},
		&PageListItemBlocks{},
		&PageListItemText{},
		&PageListOrderedItemBlocks{},
		&PageListOrderedItemText{},
		&PageRelatedArticle{},
		&PageTableCell{},
		&PageTableRow{},
		&PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{},
		&PasswordKdfAlgoUnknown{},
		&PaymentCharge{},
		&PaymentRequestedInfo{},
		&PaymentSavedCredentialsCard{},
		&PaymentsBankCardData{},
		&PaymentsClearSavedInfoParams{},
		&PaymentsGetBankCardDataParams{},
		&PaymentsGetPaymentFormParams{},
		&PaymentsGetPaymentReceiptParams{},
		&PaymentsGetSavedInfoParams{},
		&PaymentsPaymentForm{},
		&PaymentsPaymentReceipt{},
		&PaymentsPaymentResultObj{},
		&PaymentsPaymentVerificationNeeded{},
		&PaymentsSavedInfo{},
		&PaymentsSendPaymentFormParams{},
		&PaymentsValidateRequestedInfoParams{},
		&PaymentsValidatedRequestedInfo{},
		&PeerBlocked{},
		&PeerChannel{},
		&PeerChat{},
		&PeerLocatedObj{},
		&PeerNotifySettings{},
		&PeerSelfLocated{},
		&PeerSettings{},
		&PeerUser{},
		&PhoneAcceptCallParams{},
		&PhoneCallAccepted{},
		&PhoneCallDiscarded{},
		&PhoneCallEmpty{},
		&PhoneCallObj{},
		&PhoneCallProtocol{},
		&PhoneCallRequested{},
		&PhoneCallWaiting{},
		&PhoneConfirmCallParams{},
		&PhoneConnectionObj{},
		&PhoneConnectionWebrtc{},
		&PhoneDiscardCallParams{},
		&PhoneGetCallConfigParams{},
		&PhonePhoneCall{},
		&PhoneReceivedCallParams{},
		&PhoneRequestCallParams{},
		&PhoneSaveCallDebugParams{},
		&PhoneSendSignalingDataParams{},
		&PhoneSetCallRatingParams{},
		&PhotoCachedSize{},
		&PhotoEmpty{},
		&PhotoObj{},
		&PhotoPathSize{},
		&PhotoSizeEmpty{},
		&PhotoSizeObj{},
		&PhotoSizeProgressive{},
		&PhotoStrippedSize{},
		&PhotosDeletePhotosParams{},
		&PhotosGetUserPhotosParams{},
		&PhotosPhoto{},
		&PhotosPhotosObj{},
		&PhotosPhotosSlice{},
		&PhotosUpdateProfilePhotoParams{},
		&PhotosUploadProfilePhotoParams{},
		&Poll{},
		&PollAnswer{},
		&PollAnswerVoters{},
		&PollResults{},
		&PopularContact{},
		&PostAddress{},
		&PrivacyValueAllowAll{},
		&PrivacyValueAllowChatParticipants{},
		&PrivacyValueAllowContacts{},
		&PrivacyValueAllowUsers{},
		&PrivacyValueDisallowAll{},
		&PrivacyValueDisallowChatParticipants{},
		&PrivacyValueDisallowContacts{},
		&PrivacyValueDisallowUsers{},
		&ReceivedNotifyMessage{},
		&RecentMeURLChat{},
		&RecentMeURLChatInvite{},
		&RecentMeURLStickerSet{},
		&RecentMeURLUnknown{},
		&RecentMeURLUser{},
		&ReplyInlineMarkup{},
		&ReplyKeyboardForceReply{},
		&ReplyKeyboardHide{},
		&ReplyKeyboardMarkup{},
		&RestrictionReason{},
		&SavedPhoneContact{},
		&SecureCredentialsEncrypted{},
		&SecureData{},
		&SecureFileEmpty{},
		&SecureFileObj{},
		&SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000{},
		&SecurePasswordKdfAlgoSHA512{},
		&SecurePasswordKdfAlgoUnknown{},
		&SecurePlainEmail{},
		&SecurePlainPhone{},
		&SecureRequiredTypeObj{},
		&SecureRequiredTypeOneOf{},
		&SecureSecretSettings{},
		&SecureValue{},
		&SecureValueErrorData{},
		&SecureValueErrorFile{},
		&SecureValueErrorFiles{},
		&SecureValueErrorFrontSide{},
		&SecureValueErrorObj{},
		&SecureValueErrorReverseSide{},
		&SecureValueErrorSelfie{},
		&SecureValueErrorTranslationFile{},
		&SecureValueErrorTranslationFiles{},
		&SecureValueHash{},
		&SendMessageCancelAction{},
		&SendMessageChooseContactAction{},
		&SendMessageGamePlayAction{},
		&SendMessageGeoLocationAction{},
		&SendMessageRecordAudioAction{},
		&SendMessageRecordRoundAction{},
		&SendMessageRecordVideoAction{},
		&SendMessageTypingAction{},
		&SendMessageUploadAudioAction{},
		&SendMessageUploadDocumentAction{},
		&SendMessageUploadPhotoAction{},
		&SendMessageUploadRoundAction{},
		&SendMessageUploadVideoAction{},
		&ShippingOption{},
		&StatsAbsValueAndPrev{},
		&StatsBroadcastStats{},
		&StatsDateRangeDays{},
		&StatsGetBroadcastStatsParams{},
		&StatsGetMegagroupStatsParams{},
		&StatsGetMessagePublicForwardsParams{},
		&StatsGetMessageStatsParams{},
		&StatsGraphAsync{},
		&StatsGraphError{},
		&StatsGraphObj{},
		&StatsGroupTopAdmin{},
		&StatsGroupTopInviter{},
		&StatsGroupTopPoster{},
		&StatsLoadAsyncGraphParams{},
		&StatsMegagroupStats{},
		&StatsMessageStats{},
		&StatsPercentValue{},
		&StatsURL{},
		&StickerPack{},
		&StickerSet{},
		&StickerSetCoveredObj{},
		&StickerSetMultiCovered{},
		&StickersAddStickerToSetParams{},
		&StickersChangeStickerPositionParams{},
		&StickersCreateStickerSetParams{},
		&StickersRemoveStickerFromSetParams{},
		&StickersSetStickerSetThumbParams{},
		&TextAnchor{},
		&TextBold{},
		&TextConcat{},
		&TextEmail{},
		&TextEmpty{},
		&TextFixed{},
		&TextImage{},
		&TextItalic{},
		&TextMarked{},
		&TextPhone{},
		&TextPlain{},
		&TextStrike{},
		&TextSubscript{},
		&TextSuperscript{},
		&TextURL{},
		&TextUnderline{},
		&Theme{},
		&ThemeSettings{},
		&TopPeer{},
		&TopPeerCategoryPeers{},
		&URLAuthResultAccepted{},
		&URLAuthResultDefault{},
		&URLAuthResultRequest{},
		&UpdateBotCallbackQuery{},
		&UpdateBotInlineQuery{},
		&UpdateBotInlineSend{},
		&UpdateBotPrecheckoutQuery{},
		&UpdateBotShippingQuery{},
		&UpdateBotWebhookJson{},
		&UpdateBotWebhookJsonQuery{},
		&UpdateChannel{},
		&UpdateChannelAvailableMessages{},
		&UpdateChannelMessageForwards{},
		&UpdateChannelMessageViews{},
		&UpdateChannelParticipant{},
		&UpdateChannelReadMessagesContents{},
		&UpdateChannelTooLong{},
		&UpdateChannelUserTyping{},
		&UpdateChannelWebPage{},
		&UpdateChatDefaultBannedRights{},
		&UpdateChatParticipantAdd{},
		&UpdateChatParticipantAdmin{},
		&UpdateChatParticipantDelete{},
		&UpdateChatParticipants{},
		&UpdateChatUserTyping{},
		&UpdateConfig{},
		&UpdateContactsReset{},
		&UpdateDcOptions{},
		&UpdateDeleteChannelMessages{},
		&UpdateDeleteMessages{},
		&UpdateDeleteScheduledMessages{},
		&UpdateDialogFilter{},
		&UpdateDialogFilterOrder{},
		&UpdateDialogFilters{},
		&UpdateDialogPinned{},
		&UpdateDialogUnreadMark{},
		&UpdateDraftMessage{},
		&UpdateEditChannelMessage{},
		&UpdateEditMessage{},
		&UpdateEncryptedChatTyping{},
		&UpdateEncryptedMessagesRead{},
		&UpdateEncryption{},
		&UpdateFavedStickers{},
		&UpdateFolderPeers{},
		&UpdateGeoLiveViewed{},
		&UpdateInlineBotCallbackQuery{},
		&UpdateLangPack{},
		&UpdateLangPackTooLong{},
		&UpdateLoginToken{},
		&UpdateMessageID{},
		&UpdateMessagePoll{},
		&UpdateMessagePollVote{},
		&UpdateNewChannelMessage{},
		&UpdateNewEncryptedMessage{},
		&UpdateNewMessage{},
		&UpdateNewScheduledMessage{},
		&UpdateNewStickerSet{},
		&UpdateNotifySettings{},
		&UpdatePeerBlocked{},
		&UpdatePeerLocated{},
		&UpdatePeerSettings{},
		&UpdatePhoneCall{},
		&UpdatePhoneCallSignalingData{},
		&UpdatePinnedChannelMessages{},
		&UpdatePinnedDialogs{},
		&UpdatePinnedMessages{},
		&UpdatePrivacy{},
		&UpdatePtsChanged{},
		&UpdateReadChannelDiscussionInbox{},
		&UpdateReadChannelDiscussionOutbox{},
		&UpdateReadChannelInbox{},
		&UpdateReadChannelOutbox{},
		&UpdateReadFeaturedStickers{},
		&UpdateReadHistoryInbox{},
		&UpdateReadHistoryOutbox{},
		&UpdateReadMessagesContents{},
		&UpdateRecentStickers{},
		&UpdateSavedGifs{},
		&UpdateServiceNotification{},
		&UpdateShort{},
		&UpdateShortChatMessage{},
		&UpdateShortMessage{},
		&UpdateShortSentMessage{},
		&UpdateStickerSets{},
		&UpdateStickerSetsOrder{},
		&UpdateTheme{},
		&UpdateUserName{},
		&UpdateUserPhone{},
		&UpdateUserPhoto{},
		&UpdateUserStatus{},
		&UpdateUserTyping{},
		&UpdateWebPage{},
		&UpdatesChannelDifferenceEmpty{},
		&UpdatesChannelDifferenceObj{},
		&UpdatesChannelDifferenceTooLong{},
		&UpdatesCombined{},
		&UpdatesDifferenceEmpty{},
		&UpdatesDifferenceObj{},
		&UpdatesDifferenceSlice{},
		&UpdatesDifferenceTooLong{},
		&UpdatesGetChannelDifferenceParams{},
		&UpdatesGetDifferenceParams{},
		&UpdatesGetStateParams{},
		&UpdatesObj{},
		&UpdatesState{},
		&UpdatesTooLong{},
		&UploadCdnFileObj{},
		&UploadCdnFileReuploadNeeded{},
		&UploadFileCdnRedirect{},
		&UploadFileObj{},
		&UploadGetCdnFileHashesParams{},
		&UploadGetCdnFileParams{},
		&UploadGetFileHashesParams{},
		&UploadGetFileParams{},
		&UploadGetWebFileParams{},
		&UploadReuploadCdnFileParams{},
		&UploadSaveBigFilePartParams{},
		&UploadSaveFilePartParams{},
		&UploadWebFile{},
		&UserEmpty{},
		&UserFull{},
		&UserObj{},
		&UserProfilePhotoEmpty{},
		&UserProfilePhotoObj{},
		&UserStatusEmpty{},
		&UserStatusLastMonth{},
		&UserStatusLastWeek{},
		&UserStatusOffline{},
		&UserStatusOnline{},
		&UserStatusRecently{},
		&UsersGetFullUserParams{},
		&UsersGetUsersParams{},
		&UsersSetSecureValueErrorsParams{},
		&VideoSize{},
		&WallPaperNoFile{},
		&WallPaperObj{},
		&WallPaperSettings{},
		&WebAuthorization{},
		&WebDocumentNoProxy{},
		&WebDocumentObj{},
		&WebPageAttributeTheme{},
		&WebPageEmpty{},
		&WebPageNotModified{},
		&WebPageObj{},
		&WebPagePending{})

	tl.RegisterEnums(AuthCodeTypeCall,
		AuthCodeTypeFlashCall,
		AuthCodeTypeSms,
		BaseThemeArctic,
		BaseThemeClassic,
		BaseThemeDay,
		BaseThemeNight,
		BaseThemeTinted,
		InputPrivacyKeyAddedByPhone,
		InputPrivacyKeyChatInvite,
		InputPrivacyKeyForwards,
		InputPrivacyKeyPhoneCall,
		InputPrivacyKeyPhoneNumber,
		InputPrivacyKeyPhoneP2P,
		InputPrivacyKeyProfilePhoto,
		InputPrivacyKeyStatusTimestamp,
		PhoneCallDiscardReasonBusy,
		PhoneCallDiscardReasonDisconnect,
		PhoneCallDiscardReasonHangup,
		PhoneCallDiscardReasonMissed,
		PrivacyKeyAddedByPhone,
		PrivacyKeyChatInvite,
		PrivacyKeyForwards,
		PrivacyKeyPhoneCall,
		PrivacyKeyPhoneNumber,
		PrivacyKeyPhoneP2P,
		PrivacyKeyProfilePhoto,
		PrivacyKeyStatusTimestamp,
		SecureValueTypeAddress,
		SecureValueTypeBankStatement,
		SecureValueTypeDriverLicense,
		SecureValueTypeEmail,
		SecureValueTypeIdentityCard,
		SecureValueTypeInternalPassport,
		SecureValueTypePassport,
		SecureValueTypePassportRegistration,
		SecureValueTypePersonalDetails,
		SecureValueTypePhone,
		SecureValueTypeRentalAgreement,
		SecureValueTypeTemporaryRegistration,
		SecureValueTypeUtilityBill,
		StorageFileGif,
		StorageFileJpeg,
		StorageFileMov,
		StorageFileMp3,
		StorageFileMp4,
		StorageFilePartial,
		StorageFilePdf,
		StorageFilePng,
		StorageFileUnknown,
		StorageFileWebp,
		TopPeerCategoryBotsInline,
		TopPeerCategoryBotsPm,
		TopPeerCategoryChannels,
		TopPeerCategoryCorrespondents,
		TopPeerCategoryForwardChats,
		TopPeerCategoryForwardUsers,
		TopPeerCategoryGroups,
		TopPeerCategoryPhoneCalls)
}
