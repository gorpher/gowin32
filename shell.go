/*
 * Copyright (c) 2014-2017 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gowin32

import (
	"errors"
	"github.com/gorpher/gowin32/win"
	"github.com/gorpher/gowin32/wrappers"
	"image"
	"image/color"
	"syscall"
	"unsafe"
)

type SpecialFolder uint32

const (
	FolderDesktop                SpecialFolder = wrappers.CSIDL_DESKTOP
	FolderInternet               SpecialFolder = wrappers.CSIDL_INTERNET
	FolderPrograms               SpecialFolder = wrappers.CSIDL_PROGRAMS
	FolderControls               SpecialFolder = wrappers.CSIDL_CONTROLS
	FolderPrinters               SpecialFolder = wrappers.CSIDL_PRINTERS
	FolderPersonal               SpecialFolder = wrappers.CSIDL_PERSONAL
	FolderFavorites              SpecialFolder = wrappers.CSIDL_FAVORITES
	FolderStartup                SpecialFolder = wrappers.CSIDL_STARTUP
	FolderRecent                 SpecialFolder = wrappers.CSIDL_RECENT
	FolderSendTo                 SpecialFolder = wrappers.CSIDL_SENDTO
	FolderBitBucket              SpecialFolder = wrappers.CSIDL_BITBUCKET
	FolderStartMenu              SpecialFolder = wrappers.CSIDL_STARTMENU
	FolderMyDocuments            SpecialFolder = wrappers.CSIDL_MYDOCUMENTS
	FolderMyMusic                SpecialFolder = wrappers.CSIDL_MYMUSIC
	FolderMyVideo                SpecialFolder = wrappers.CSIDL_MYVIDEO
	FolderDesktopDirectory       SpecialFolder = wrappers.CSIDL_DESKTOPDIRECTORY
	FolderDrives                 SpecialFolder = wrappers.CSIDL_DRIVES
	FolderNetwork                SpecialFolder = wrappers.CSIDL_NETWORK
	FolderNetHood                SpecialFolder = wrappers.CSIDL_NETHOOD
	FolderFonts                  SpecialFolder = wrappers.CSIDL_FONTS
	FolderTemplates              SpecialFolder = wrappers.CSIDL_TEMPLATES
	FolderCommonStartMenu        SpecialFolder = wrappers.CSIDL_COMMON_STARTMENU
	FolderCommonPrograms         SpecialFolder = wrappers.CSIDL_COMMON_PROGRAMS
	FolderCommonStartup          SpecialFolder = wrappers.CSIDL_COMMON_STARTUP
	FolderCommonDesktopDirectory SpecialFolder = wrappers.CSIDL_COMMON_DESKTOPDIRECTORY
	FolderAppData                SpecialFolder = wrappers.CSIDL_APPDATA
	FolderPrintHood              SpecialFolder = wrappers.CSIDL_PRINTHOOD
	FolderLocalAppData           SpecialFolder = wrappers.CSIDL_LOCAL_APPDATA
	FolderAltStartup             SpecialFolder = wrappers.CSIDL_ALTSTARTUP
	FolderCommonAltStartup       SpecialFolder = wrappers.CSIDL_COMMON_ALTSTARTUP
	FolderCommonFavorites        SpecialFolder = wrappers.CSIDL_COMMON_FAVORITES
	FolderInternetCache          SpecialFolder = wrappers.CSIDL_INTERNET_CACHE
	FolderCookies                SpecialFolder = wrappers.CSIDL_COOKIES
	FolderHistory                SpecialFolder = wrappers.CSIDL_HISTORY
	FolderCommonAppData          SpecialFolder = wrappers.CSIDL_COMMON_APPDATA
	FolderWindows                SpecialFolder = wrappers.CSIDL_WINDOWS
	FolderSystem                 SpecialFolder = wrappers.CSIDL_SYSTEM
	FolderProgramFiles           SpecialFolder = wrappers.CSIDL_PROGRAM_FILES
	FolderMyPictures             SpecialFolder = wrappers.CSIDL_MYPICTURES
	FolderProfile                SpecialFolder = wrappers.CSIDL_PROFILE
	FolderSystemX86              SpecialFolder = wrappers.CSIDL_SYSTEMX86
	FolderProgramFilesX86        SpecialFolder = wrappers.CSIDL_PROGRAM_FILESX86
	FolderProgramFilesCommon     SpecialFolder = wrappers.CSIDL_PROGRAM_FILES_COMMON
	FolderProgramFilesCommonX86  SpecialFolder = wrappers.CSIDL_PROGRAM_FILES_COMMONX86
	FolderCommonTemplates        SpecialFolder = wrappers.CSIDL_COMMON_TEMPLATES
	FolderCommonDocuments        SpecialFolder = wrappers.CSIDL_COMMON_DOCUMENTS
	FolderCommonAdminTools       SpecialFolder = wrappers.CSIDL_COMMON_ADMINTOOLS
	FolderAdminTools             SpecialFolder = wrappers.CSIDL_ADMINTOOLS
	FolderConnections            SpecialFolder = wrappers.CSIDL_CONNECTIONS
	FolderCommonMusic            SpecialFolder = wrappers.CSIDL_COMMON_MUSIC
	FolderCommonPictures         SpecialFolder = wrappers.CSIDL_COMMON_PICTURES
	FolderCommonVideo            SpecialFolder = wrappers.CSIDL_COMMON_VIDEO
	FolderResources              SpecialFolder = wrappers.CSIDL_RESOURCES
	FolderResourcesLocalized     SpecialFolder = wrappers.CSIDL_RESOURCES_LOCALIZED
	FolderCommonOEMLinks         SpecialFolder = wrappers.CSIDL_COMMON_OEM_LINKS
	FolderCDBurnArea             SpecialFolder = wrappers.CSIDL_CDBURN_AREA
	FolderComputersNearMe        SpecialFolder = wrappers.CSIDL_COMPUTERSNEARME
)

type KnownFolder wrappers.GUID

var (
	KnownFolderNetworkFolder          = KnownFolder(wrappers.FOLDERID_NetworkFolder)
	KnownFolderComputerFolder         = KnownFolder(wrappers.FOLDERID_ComputerFolder)
	KnownFolderInternetFolder         = KnownFolder(wrappers.FOLDERID_InternetFolder)
	KnownFolderControlPanelFolder     = KnownFolder(wrappers.FOLDERID_ControlPanelFolder)
	KnownFolderPrintersFolder         = KnownFolder(wrappers.FOLDERID_PrintersFolder)
	KnownFolderSyncManagerFolder      = KnownFolder(wrappers.FOLDERID_SyncManagerFolder)
	KnownFolderSyncSetupFolder        = KnownFolder(wrappers.FOLDERID_SyncSetupFolder)
	KnownFolderConflictFolder         = KnownFolder(wrappers.FOLDERID_ConflictFolder)
	KnownFolderSyncResultsFolder      = KnownFolder(wrappers.FOLDERID_SyncResultsFolder)
	KnownFolderRecycleBinFolder       = KnownFolder(wrappers.FOLDERID_RecycleBinFolder)
	KnownFolderConnectionsFolder      = KnownFolder(wrappers.FOLDERID_ConnectionsFolder)
	KnownFolderFonts                  = KnownFolder(wrappers.FOLDERID_Fonts)
	KnownFolderDesktop                = KnownFolder(wrappers.FOLDERID_Desktop)
	KnownFolderStartup                = KnownFolder(wrappers.FOLDERID_Startup)
	KnownFolderPrograms               = KnownFolder(wrappers.FOLDERID_Programs)
	KnownFolderStartMenu              = KnownFolder(wrappers.FOLDERID_StartMenu)
	KnownFolderRecent                 = KnownFolder(wrappers.FOLDERID_Recent)
	KnownFolderSendTo                 = KnownFolder(wrappers.FOLDERID_SendTo)
	KnownFolderDocuments              = KnownFolder(wrappers.FOLDERID_Documents)
	KnownFolderFavorites              = KnownFolder(wrappers.FOLDERID_Favorites)
	KnownFolderNetHood                = KnownFolder(wrappers.FOLDERID_NetHood)
	KnownFolderPrintHood              = KnownFolder(wrappers.FOLDERID_PrintHood)
	KnownFolderTemplates              = KnownFolder(wrappers.FOLDERID_Templates)
	KnownFolderCommonStartup          = KnownFolder(wrappers.FOLDERID_CommonStartup)
	KnownFolderCommonPrograms         = KnownFolder(wrappers.FOLDERID_CommonPrograms)
	KnownFolderCommonStartMenu        = KnownFolder(wrappers.FOLDERID_CommonStartMenu)
	KnownFolderPublicDesktop          = KnownFolder(wrappers.FOLDERID_PublicDesktop)
	KnownFolderProgramData            = KnownFolder(wrappers.FOLDERID_ProgramData)
	KnownFolderCommonTemplates        = KnownFolder(wrappers.FOLDERID_CommonTemplates)
	KnownFolderPublicDocuments        = KnownFolder(wrappers.FOLDERID_PublicDocuments)
	KnownFolderRoamingAppData         = KnownFolder(wrappers.FOLDERID_RoamingAppData)
	KnownFolderLocalAppData           = KnownFolder(wrappers.FOLDERID_LocalAppData)
	KnownFolderLocalAppDataLow        = KnownFolder(wrappers.FOLDERID_LocalAppDataLow)
	KnownFolderInternetCache          = KnownFolder(wrappers.FOLDERID_InternetCache)
	KnownFolderCookies                = KnownFolder(wrappers.FOLDERID_Cookies)
	KnownFolderHistory                = KnownFolder(wrappers.FOLDERID_History)
	KnownFolderSystem                 = KnownFolder(wrappers.FOLDERID_System)
	KnownFolderSystemX86              = KnownFolder(wrappers.FOLDERID_SystemX86)
	KnownFolderWindows                = KnownFolder(wrappers.FOLDERID_Windows)
	KnownFolderProfile                = KnownFolder(wrappers.FOLDERID_Profile)
	KnownFolderPictures               = KnownFolder(wrappers.FOLDERID_Pictures)
	KnownFolderProgramFilesX86        = KnownFolder(wrappers.FOLDERID_ProgramFilesX86)
	KnownFolderProgramFilesCommonX86  = KnownFolder(wrappers.FOLDERID_ProgramFilesCommonX86)
	KnownFolderProgramFilesX64        = KnownFolder(wrappers.FOLDERID_ProgramFilesX64)
	KnownFolderProgramFilesCommonX64  = KnownFolder(wrappers.FOLDERID_ProgramFilesCommonX64)
	KnownFolderProgramFiles           = KnownFolder(wrappers.FOLDERID_ProgramFiles)
	KnownFolderProgramFilesCommon     = KnownFolder(wrappers.FOLDERID_ProgramFilesCommon)
	KnownFolderUserProgramFiles       = KnownFolder(wrappers.FOLDERID_UserProgramFiles)
	KnownFolderUserProgramFilesCommon = KnownFolder(wrappers.FOLDERID_UserProgramFilesCommon)
	KnownFolderAdminTools             = KnownFolder(wrappers.FOLDERID_AdminTools)
	KnownFolderCommonAdminTools       = KnownFolder(wrappers.FOLDERID_CommonAdminTools)
	KnownFolderMusic                  = KnownFolder(wrappers.FOLDERID_Music)
	KnownFolderVideos                 = KnownFolder(wrappers.FOLDERID_Videos)
	KnownFolderRingtones              = KnownFolder(wrappers.FOLDERID_Ringtones)
	KnownFolderPublicPictures         = KnownFolder(wrappers.FOLDERID_PublicPictures)
	KnownFolderPublicMusic            = KnownFolder(wrappers.FOLDERID_PublicMusic)
	KnownFolderPublicVideos           = KnownFolder(wrappers.FOLDERID_PublicVideos)
	KnownFolderPublicRingtones        = KnownFolder(wrappers.FOLDERID_PublicRingtones)
	KnownFolderResourceDir            = KnownFolder(wrappers.FOLDERID_ResourceDir)
	KnownFolderLocalizedResourcesDir  = KnownFolder(wrappers.FOLDERID_LocalizedResourcesDir)
	KnownFolderCommonOEMLinks         = KnownFolder(wrappers.FOLDERID_CommonOEMLinks)
	KnownFolderCDBurning              = KnownFolder(wrappers.FOLDERID_CDBurning)
	KnownFolderUserProfiles           = KnownFolder(wrappers.FOLDERID_UserProfiles)
	KnownFolderPlaylists              = KnownFolder(wrappers.FOLDERID_Playlists)
	KnownFolderSamplePlaylists        = KnownFolder(wrappers.FOLDERID_SamplePlaylists)
	KnownFolderSampleMusic            = KnownFolder(wrappers.FOLDERID_SampleMusic)
	KnownFolderSamplePictures         = KnownFolder(wrappers.FOLDERID_SamplePictures)
	KnownFolderSampleVideos           = KnownFolder(wrappers.FOLDERID_SampleVideos)
	KnownFolderPhotoAlbums            = KnownFolder(wrappers.FOLDERID_PhotoAlbums)
	KnownFolderPublic                 = KnownFolder(wrappers.FOLDERID_Public)
	KnownFolderChangeRemovePrograms   = KnownFolder(wrappers.FOLDERID_ChangeRemovePrograms)
	KnownFolderAppUpdates             = KnownFolder(wrappers.FOLDERID_AppUpdates)
	KnownFolderAddNewPrograms         = KnownFolder(wrappers.FOLDERID_AddNewPrograms)
	KnownFolderDownloads              = KnownFolder(wrappers.FOLDERID_Downloads)
	KnownFolderPublicDownloads        = KnownFolder(wrappers.FOLDERID_PublicDownloads)
	KnownFolderSavedSearches          = KnownFolder(wrappers.FOLDERID_SavedSearches)
	KnownFolderQuickLaunch            = KnownFolder(wrappers.FOLDERID_QuickLaunch)
	KnownFolderContacts               = KnownFolder(wrappers.FOLDERID_Contacts)
	KnownFolderSidebarParts           = KnownFolder(wrappers.FOLDERID_SidebarParts)
	KnownFolderSidebarDefaultParts    = KnownFolder(wrappers.FOLDERID_SidebarDefaultParts)
	KnownFolderPublicGameTasks        = KnownFolder(wrappers.FOLDERID_PublicGameTasks)
	KnownFolderGameTasks              = KnownFolder(wrappers.FOLDERID_GameTasks)
	KnownFolderSavedGames             = KnownFolder(wrappers.FOLDERID_SavedGames)
	KnownFolderGames                  = KnownFolder(wrappers.FOLDERID_Games)
	KnownFolderSearchMAPI             = KnownFolder(wrappers.FOLDERID_SEARCH_MAPI)
	KnownFolderSearchCSC              = KnownFolder(wrappers.FOLDERID_SEARCH_CSC)
	KnownFolderLinks                  = KnownFolder(wrappers.FOLDERID_Links)
	KnownFolderUserLinks              = KnownFolder(wrappers.FOLDERID_UserLinks)
	KnownFolderUserLibraries          = KnownFolder(wrappers.FOLDERID_UserLibraries)
	KnownFolderSearchHome             = KnownFolder(wrappers.FOLDERID_SearchHome)
	KnownFolderOriginalImages         = KnownFolder(wrappers.FOLDERID_OriginalImages)
	KnownFolderDocumentsLibrary       = KnownFolder(wrappers.FOLDERID_DocumentsLibrary)
	KnownFolderMusicLibrary           = KnownFolder(wrappers.FOLDERID_MusicLibrary)
	KnownFolderPicturesLibrary        = KnownFolder(wrappers.FOLDERID_PicturesLibrary)
	KnownFolderVideosLibrary          = KnownFolder(wrappers.FOLDERID_VideosLibrary)
	KnownFolderRecordedTVLibrary      = KnownFolder(wrappers.FOLDERID_RecordedTVLibrary)
	KnownFolderHomeGroup              = KnownFolder(wrappers.FOLDERID_HomeGroup)
	KnownFolderDeviceMetadataStore    = KnownFolder(wrappers.FOLDERID_DeviceMetadataStore)
	KnownFolderLibraries              = KnownFolder(wrappers.FOLDERID_Libraries)
	KnownFolderPublicLibraries        = KnownFolder(wrappers.FOLDERID_PublicLibraries)
	KnownFolderUserPinned             = KnownFolder(wrappers.FOLDERID_UserPinned)
	KnownFolderImplicitAppShortcuts   = KnownFolder(wrappers.FOLDERID_ImplicitAppShortcuts)
)

func GetSpecialFolderPath(folder SpecialFolder) (string, error) {
	buf := [wrappers.MAX_PATH]uint16{}
	if hr := wrappers.SHGetFolderPath(0, uint32(folder), 0, 0, &buf[0]); wrappers.FAILED(hr) {
		return "", NewWindowsError("SHGetFolderPath", COMError(hr))
	}
	return syscall.UTF16ToString((&buf)[:]), nil
}

func GetKnownFolderPath(folder KnownFolder) (string, error) {
	var path *uint16
	if hr := wrappers.SHGetKnownFolderPath((*wrappers.GUID)(&folder), 0, 0, &path); wrappers.FAILED(hr) {
		return "", NewWindowsError("SHGetKnownFolderPath", COMError(hr))
	}
	defer wrappers.CoTaskMemFree((*byte)(unsafe.Pointer(path)))
	return LpstrToString(path), nil
}

func ShellDelete(fileSpec string) error {
	return wrappers.SHFileOperation(&wrappers.SHFILEOPSTRUCT{
		Func:  wrappers.FO_DELETE,
		From:  Lpcwstr(fileSpec),
		Flags: wrappers.FOF_NO_UI,
	})
}

func ShellCopy(source string, destination string) error {
	return wrappers.SHFileOperation(&wrappers.SHFILEOPSTRUCT{
		Func:  wrappers.FO_COPY,
		From:  Lpcwstr(source),
		To:    Lpcwstr(destination),
		Flags: wrappers.FOF_NO_UI,
	})
}

func ExtractIconToImageByExt(filename string) (image.Image, error) {
	var shFile wrappers.SHFILEINFO
	err := wrappers.SHGetFileInfoW(
		Lpcwstr(filename),
		0,
		uintptr(unsafe.Pointer(&shFile)),
		uint32(unsafe.Sizeof(shFile)),
		wrappers.SHGFI_ICON|wrappers.SHGFI_USEFILEATTRIBUTES)
	if err != nil {
		return nil, err
	}
	defer win.DestroyIcon(win.HICON(shFile.HIcon))
	return hICONTOImage(shFile.HIcon)
}

func ExtractIconToImage(filename string) (image.Image, error) {
	large := []syscall.Handle{0}
	err := wrappers.ExtractIconExW(Lpcwstr(filename), 0, &large[0], nil, 1)
	if err != nil {
		return nil, err
	}
	defer win.DestroyIcon(win.HICON(large[0]))
	return hICONTOImage(large[0])
}

// ExtractPrivateExtractIcons 提取exe高清图标
func ExtractPrivateExtractIcons(filename string, w, h uint32) (image.Image, error) {
	large := []syscall.Handle{0}
	var piconId uint32 = 0
	err := wrappers.PrivateExtractIcons(filename, 0, w, h, &large[0], &piconId, 1, 0)
	if err != nil {
		return nil, err
	}
	if large == nil {
		return nil, errors.New("can't extract  icon")
	}
	defer win.DestroyIcon(win.HICON(large[0]))
	return hICONTOImage(large[0])
}

func hICONTOImage(handle syscall.Handle) (image.Image, error) {
	var iconInfo win.ICONINFO
	err := wrappers.GetIconInfo(handle, uintptr(unsafe.Pointer(&iconInfo)))
	if err != nil {
		return nil, err
	}

	w := int32(iconInfo.XHotspot * 2)
	h := int32(iconInfo.YHotspot * 2)

	var bitmapInfo win.BITMAPINFOHEADER
	bitmapInfo.BiSize = uint32(unsafe.Sizeof(bitmapInfo))
	bitmapInfo.BiWidth = w
	bitmapInfo.BiHeight = -h
	bitmapInfo.BiPlanes = 1
	bitmapInfo.BiBitCount = 32
	bitmapInfo.BiCompression = win.BI_RGB
	bitmapInfo.BiSizeImage = uint32(w * h * 4)
	bitmapInfo.BiXPelsPerMeter = 0
	bitmapInfo.BiYPelsPerMeter = 0
	bitmapInfo.BiClrUsed = 0
	bitmapInfo.BiClrImportant = 0
	dc := win.GetDC(0)
	hdc := win.CreateCompatibleDC(dc)
	ok := win.ReleaseDC(0, hdc)
	if ok {
		return nil, errors.New("ReleaseDC Error")
	}
	defer win.DeleteDC(hdc)
	var bits unsafe.Pointer
	winBitmap := win.CreateDIBSection(hdc, &bitmapInfo, win.DIB_RGB_COLORS, &bits, 0, 0)
	defer win.DeleteObject(win.HGDIOBJ(winBitmap))

	var pixels = (*[1 << 30]byte)(bits)[0:bitmapInfo.BiSizeImage]
	win.SelectObject(hdc, win.HGDIOBJ(winBitmap))
	success := win.DrawIconEx(hdc, 0, 0, win.HICON(handle), w, h, 0, 0, win.DI_NORMAL)
	if ok {
		return nil, errors.New("DrawIcon Error")
	}
	hasAlpha := false
	rgba := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(w),
			Y: int(h),
		},
	})
	for y := int32(0); y < h; y++ {
		for x := int32(0); x < w; x++ {
			if pixels[((y*w+x)*4)+3] > 0 {
				hasAlpha = true
			}

			rgba.SetRGBA(int(x), int(y), color.RGBA{
				A: pixels[((y*w+x)*4)+3],
				R: pixels[((y*w+x)*4)+2],
				G: pixels[((y*w+x)*4)+1],
				B: pixels[((y*w+x)*4)+0],
			})
		}
	}
	if hasAlpha {
		return rgba, nil
	}
	success = win.DrawIconEx(hdc, 0, 0, win.HICON(syscall.Handle(handle)), w, h, 0, 0, win.DI_MASK)
	if success {
		for y := int32(0); y < h; y++ {
			for x := int32(0); x < w; x++ {
				tmp := rgba.RGBAAt(int(x), int(y))
				if (pixels[((y*w+x)*4)+2] | pixels[((y*w+x)*4)+1] | pixels[((y*w+x)*4)+0]) == 0 {
					tmp.A = 0xFF
					rgba.SetRGBA(int(x), int(y), tmp)
				}
			}
		}
	}
	return rgba, nil
}
