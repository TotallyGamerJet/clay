package sdl3

import "github.com/ebitengine/purego/objc"

func init() {
	// Note: on Mac OSX, SDL2 for some reason decides to automatically disable momentum scrolling on macbook trackpads.
	// We reenable here
	key := objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"), "AppleMomentumScrollSupported")
	objc.ID(objc.GetClass("NSUserDefaults")).Send(objc.RegisterName("standardUserDefaults")).
		Send(objc.RegisterName("setBool:forKey:"), true, key)
}
