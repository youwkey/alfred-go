package alfred

import (
	"os"
	"testing"
)

func TestIcons(t *testing.T) {
	icons := []*Icon{
		IconARDocument,
		IconARObject,
		IconAccounts,
		IconActions,
		IconAirDrop,
		IconAlertCautionBadge,
		IconAlertNote,
		IconAlertStop,
		IconAliasBadge,
		IconAllMyFiles,
		IconApplicationsFolder,
		IconBackwardArrow,
		IconBonjour,
		IconBookmark,
		IconBurnableFolder,
		IconBurning,
		IconCDAudioVolume,
		IconClippingPicture,
		IconClippingSound,
		IconClippingText,
		IconClippingUnknown,
		IconClock,
		IconColorSyncProfile,
		IconConnectTo,
		IconDesktopFolder,
		IconDeveloperFolder,
		IconDocumentsFolder,
		IconDownloadsFolder,
		IconDropFolderBadge,
		IconEjectMedia,
		IconErasing,
		IconEveryone,
		IconExecutableBinary,
		IconFavoriteItems,
		IconFileVault,
		IconFinder,
		IconForwardArrow,
		IconFullTrash,
		IconGeneral,
		IconGenericAirDisk,
		IconGenericApplication,
		IconGenericDocument,
		IconGenericFileServer,
		IconGenericFolder,
		IconGenericFont,
		IconGenericNetwork,
		IconGenericQuestionMark,
		IconGenericSharepoint,
		IconGenericSpeaker,
		IconGenericStationery,
		IconGenericTimeMachineDisk,
		IconGenericURL,
		IconGenericWindow,
		IconGrid,
		IconGroupFolder,
		IconGroup,
		IconGuestUser,
		IconHelp,
		IconHomeFolder,
		IconInternetLocation,
		IconKEXT,
		IconKeepArranged,
		IconLibraryFolder,
		IconLockedBadge,
		IconLocked,
		IconMagnifyingGlass,
		IconMovieFolder,
		IconMultipleItems,
		IconMusicFolder,
		IconNetBootVolume,
		IconNewFolderBadge,
		IconNoWrite,
		IconNotLoaded,
		IconNotifications,
		IconOpenFolder,
		IconPicturesFolder,
		IconPrivateFolderBadge,
		IconProblemReport,
		IconProfileBackgroundColor,
		IconProfileFont,
		IconProfileFontAndColor,
		IconPublicFolder,
		IconReadOnlyFolderBadge,
		IconRealityFile,
		IconRecentItems,
		IconRightContainerArrow,
		IconServerApplicationsFolder,
		IconSidebarAirDrop,
		IconSidebarAirportDisk,
		IconSidebarAirportExpress,
		IconSidebarAirportExtreme,
		IconSidebarAirportExtremeTower,
		IconSidebarAllMyFiles,
		IconSidebarApplicationsFolder,
		IconSidebarBonjour,
		IconSidebarBurnFolder,
		IconSidebarDesktopFolder,
		IconSidebarDisplay,
		IconSidebarDocumentsFolder,
		IconSidebarDownloadsFolder,
		IconSidebarDropBoxFolder,
		IconSidebarExternalDisk,
		IconSidebarGenericFile,
		IconSidebarGenericFolder,
		IconSidebarHomeFolder,
		IconSidebarInternalDisk,
		IconSidebarLaptop,
		IconSidebarMacMini,
		IconSidebarMacPro,
		IconSidebarMacProCylinder,
		IconSidebarMoviesFolder,
		IconSidebarMusicFolder,
		IconSidebarNetwork,
		IconSidebarOpticalDisk,
		IconSidebarPC,
		IconSidebarPicturesFolder,
		IconSidebarPrefs,
		IconSidebarRecents,
		IconSidebarRemovableDisk,
		IconSidebarServerDrive,
		IconSidebarSmartFolder,
		IconSidebarTimeCapsule,
		IconSidebarTimeMachine,
		IconSidebarUtilitiesFolder,
		IconSidebarXserve,
		IconSidebarICloud,
		IconSidebarIDisk,
		IconSidebarIMac,
		IconSidebarIPad,
		IconSidebarIPhone,
		IconSidebarIPodTouch,
		IconSitesFolder,
		IconSmartFolder,
		IconSync,
		IconSystemFolder,
		IconToolbarAdvanced,
		IconToolbarCustomize,
		IconToolbarDelete,
		IconToolbarFavorites,
		IconToolbarInfo,
		IconToolbarLabels,
		IconTrash,
		IconUnknownFSObject,
		IconUnlocked,
		IconUnsupported,
		IconUser,
		IconUserUnknown,
		IconUsersFolder,
		IconUtilitiesFolder,
		IconVCard,
		IconIDiskGeneric,
		IconIDiskUser,
		IconPublicGenericLCD,
		IconPublicGenericPC,
	}

	for _, icon := range icons {
		t.Run(icon.Path, func(t *testing.T) {
			if _, err := os.Stat(icon.Path); err != nil {
				t.Errorf("icon %v not exists", icon.Path)
			}
		})
	}
}
