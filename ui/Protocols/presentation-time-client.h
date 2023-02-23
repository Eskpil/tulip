/* Generated by wayland-scanner 1.21.90 */

#ifndef PRESENTATION_TIME_CLIENT_PROTOCOL_H
#define PRESENTATION_TIME_CLIENT_PROTOCOL_H

#include <stdint.h>
#include <stddef.h>
#include "wayland-client.h"

#ifdef  __cplusplus
extern "C" {
#endif

/**
 * @page page_presentation_time The presentation_time protocol
 * @section page_ifaces_presentation_time Interfaces
 * - @subpage page_iface_wp_presentation - timed presentation related wl_surface requests
 * - @subpage page_iface_wp_presentation_feedback - presentation time feedback event
 * @section page_copyright_presentation_time Copyright
 * <pre>
 *
 * Copyright © 2013-2014 Collabora, Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice (including the next
 * paragraph) shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 * </pre>
 */
struct wl_output;
struct wl_surface;
struct wp_presentation;
struct wp_presentation_feedback;

#ifndef WP_PRESENTATION_INTERFACE
#define WP_PRESENTATION_INTERFACE
/**
 * @page page_iface_wp_presentation wp_presentation
 * @section page_iface_wp_presentation_desc Description
 *
 *
 *
 *
 * The main feature of this interface is accurate presentation
 * timing feedback to ensure smooth video playback while maintaining
 * audio/video synchronization. Some features use the concept of a
 * presentation clock, which is defined in the
 * presentation.clock_id event.
 *
 * A content update for a wl_surface is submitted by a
 * wl_surface.commit request. Request 'feedback' associates with
 * the wl_surface.commit and provides feedback on the content
 * update, particularly the final realized presentation time.
 *
 *
 *
 * When the final realized presentation time is available, e.g.
 * after a framebuffer flip completes, the requested
 * presentation_feedback.presented events are sent. The final
 * presentation time can differ from the compositor's predicted
 * display update time and the update's target time, especially
 * when the compositor misses its target vertical blanking period.
 * @section page_iface_wp_presentation_api API
 * See @ref iface_wp_presentation.
 */
/**
 * @defgroup iface_wp_presentation The wp_presentation interface
 *
 *
 *
 *
 * The main feature of this interface is accurate presentation
 * timing feedback to ensure smooth video playback while maintaining
 * audio/video synchronization. Some features use the concept of a
 * presentation clock, which is defined in the
 * presentation.clock_id event.
 *
 * A content update for a wl_surface is submitted by a
 * wl_surface.commit request. Request 'feedback' associates with
 * the wl_surface.commit and provides feedback on the content
 * update, particularly the final realized presentation time.
 *
 *
 *
 * When the final realized presentation time is available, e.g.
 * after a framebuffer flip completes, the requested
 * presentation_feedback.presented events are sent. The final
 * presentation time can differ from the compositor's predicted
 * display update time and the update's target time, especially
 * when the compositor misses its target vertical blanking period.
 */
extern const struct wl_interface wp_presentation_interface;
#endif
#ifndef WP_PRESENTATION_FEEDBACK_INTERFACE
#define WP_PRESENTATION_FEEDBACK_INTERFACE
/**
 * @page page_iface_wp_presentation_feedback wp_presentation_feedback
 * @section page_iface_wp_presentation_feedback_desc Description
 *
 * A presentation_feedback object returns an indication that a
 * wl_surface content update has become visible to the user.
 * One object corresponds to one content update submission
 * (wl_surface.commit). There are two possible outcomes: the
 * content update is presented to the user, and a presentation
 * timestamp delivered; or, the user did not see the content
 * update because it was superseded or its surface destroyed,
 * and the content update is discarded.
 *
 * Once a presentation_feedback object has delivered a 'presented'
 * or 'discarded' event it is automatically destroyed.
 * @section page_iface_wp_presentation_feedback_api API
 * See @ref iface_wp_presentation_feedback.
 */
/**
 * @defgroup iface_wp_presentation_feedback The wp_presentation_feedback interface
 *
 * A presentation_feedback object returns an indication that a
 * wl_surface content update has become visible to the user.
 * One object corresponds to one content update submission
 * (wl_surface.commit). There are two possible outcomes: the
 * content update is presented to the user, and a presentation
 * timestamp delivered; or, the user did not see the content
 * update because it was superseded or its surface destroyed,
 * and the content update is discarded.
 *
 * Once a presentation_feedback object has delivered a 'presented'
 * or 'discarded' event it is automatically destroyed.
 */
extern const struct wl_interface wp_presentation_feedback_interface;
#endif

#ifndef WP_PRESENTATION_ERROR_ENUM
#define WP_PRESENTATION_ERROR_ENUM
/**
 * @ingroup iface_wp_presentation
 * fatal presentation errors
 *
 * These fatal protocol errors may be emitted in response to
 * illegal presentation requests.
 */
enum wp_presentation_error {
	/**
	 * invalid value in tv_nsec
	 */
	WP_PRESENTATION_ERROR_INVALID_TIMESTAMP = 0,
	/**
	 * invalid flag
	 */
	WP_PRESENTATION_ERROR_INVALID_FLAG = 1,
};
#endif /* WP_PRESENTATION_ERROR_ENUM */

/**
 * @ingroup iface_wp_presentation
 * @struct wp_presentation_listener
 */
struct wp_presentation_listener {
	/**
	 * clock ID for timestamps
	 *
	 * This event tells the client in which clock domain the
	 * compositor interprets the timestamps used by the presentation
	 * extension. This clock is called the presentation clock.
	 *
	 * The compositor sends this event when the client binds to the
	 * presentation interface. The presentation clock does not change
	 * during the lifetime of the client connection.
	 *
	 * The clock identifier is platform dependent. On Linux/glibc, the
	 * identifier value is one of the clockid_t values accepted by
	 * clock_gettime(). clock_gettime() is defined by POSIX.1-2001.
	 *
	 * Timestamps in this clock domain are expressed as tv_sec_hi,
	 * tv_sec_lo, tv_nsec triples, each component being an unsigned
	 * 32-bit value. Whole seconds are in tv_sec which is a 64-bit
	 * value combined from tv_sec_hi and tv_sec_lo, and the additional
	 * fractional part in tv_nsec as nanoseconds. Hence, for valid
	 * timestamps tv_nsec must be in [0, 999999999].
	 *
	 * Note that clock_id applies only to the presentation clock, and
	 * implies nothing about e.g. the timestamps used in the Wayland
	 * core protocol input events.
	 *
	 * Compositors should prefer a clock which does not jump and is not
	 * slewed e.g. by NTP. The absolute value of the clock is
	 * irrelevant. Precision of one millisecond or better is
	 * recommended. Clients must be able to query the current clock
	 * value directly, not by asking the compositor.
	 * @param clk_id platform clock identifier
	 */
	void (*clock_id)(void *data,
			 struct wp_presentation *wp_presentation,
			 uint32_t clk_id);
};

/**
 * @ingroup iface_wp_presentation
 */
static inline int
wp_presentation_add_listener(struct wp_presentation *wp_presentation,
			     const struct wp_presentation_listener *listener, void *data)
{
	return wl_proxy_add_listener((struct wl_proxy *) wp_presentation,
				     (void (**)(void)) listener, data);
}

#define WP_PRESENTATION_DESTROY 0
#define WP_PRESENTATION_FEEDBACK 1

/**
 * @ingroup iface_wp_presentation
 */
#define WP_PRESENTATION_CLOCK_ID_SINCE_VERSION 1

/**
 * @ingroup iface_wp_presentation
 */
#define WP_PRESENTATION_DESTROY_SINCE_VERSION 1
/**
 * @ingroup iface_wp_presentation
 */
#define WP_PRESENTATION_FEEDBACK_SINCE_VERSION 1

/** @ingroup iface_wp_presentation */
static inline void
wp_presentation_set_user_data(struct wp_presentation *wp_presentation, void *user_data)
{
	wl_proxy_set_user_data((struct wl_proxy *) wp_presentation, user_data);
}

/** @ingroup iface_wp_presentation */
static inline void *
wp_presentation_get_user_data(struct wp_presentation *wp_presentation)
{
	return wl_proxy_get_user_data((struct wl_proxy *) wp_presentation);
}

static inline uint32_t
wp_presentation_get_version(struct wp_presentation *wp_presentation)
{
	return wl_proxy_get_version((struct wl_proxy *) wp_presentation);
}

/**
 * @ingroup iface_wp_presentation
 *
 * Informs the server that the client will no longer be using
 * this protocol object. Existing objects created by this object
 * are not affected.
 */
static inline void
wp_presentation_destroy(struct wp_presentation *wp_presentation)
{
	wl_proxy_marshal_flags((struct wl_proxy *) wp_presentation,
			 WP_PRESENTATION_DESTROY, NULL, wl_proxy_get_version((struct wl_proxy *) wp_presentation), WL_MARSHAL_FLAG_DESTROY);
}

/**
 * @ingroup iface_wp_presentation
 *
 * Request presentation feedback for the current content submission
 * on the given surface. This creates a new presentation_feedback
 * object, which will deliver the feedback information once. If
 * multiple presentation_feedback objects are created for the same
 * submission, they will all deliver the same information.
 *
 * For details on what information is returned, see the
 * presentation_feedback interface.
 */
static inline struct wp_presentation_feedback *
wp_presentation_feedback(struct wp_presentation *wp_presentation, struct wl_surface *surface)
{
	struct wl_proxy *callback;

	callback = wl_proxy_marshal_flags((struct wl_proxy *) wp_presentation,
			 WP_PRESENTATION_FEEDBACK, &wp_presentation_feedback_interface, wl_proxy_get_version((struct wl_proxy *) wp_presentation), 0, surface, NULL);

	return (struct wp_presentation_feedback *) callback;
}

#ifndef WP_PRESENTATION_FEEDBACK_KIND_ENUM
#define WP_PRESENTATION_FEEDBACK_KIND_ENUM
/**
 * @ingroup iface_wp_presentation_feedback
 * bitmask of flags in presented event
 *
 * These flags provide information about how the presentation of
 * the related content update was done. The intent is to help
 * clients assess the reliability of the feedback and the visual
 * quality with respect to possible tearing and timings.
 */
enum wp_presentation_feedback_kind {
	/**
	 * presentation was vsync'd
	 *
	 * The presentation was synchronized to the "vertical retrace" by
	 * the display hardware such that tearing does not happen. Relying
	 * on software scheduling is not acceptable for this flag. If
	 * presentation is done by a copy to the active frontbuffer, then
	 * it must guarantee that tearing cannot happen.
	 */
	WP_PRESENTATION_FEEDBACK_KIND_VSYNC = 0x1,
	/**
	 * hardware provided the presentation timestamp
	 *
	 * The display hardware provided measurements that the hardware
	 * driver converted into a presentation timestamp. Sampling a clock
	 * in software is not acceptable for this flag.
	 */
	WP_PRESENTATION_FEEDBACK_KIND_HW_CLOCK = 0x2,
	/**
	 * hardware signalled the start of the presentation
	 *
	 * The display hardware signalled that it started using the new
	 * image content. The opposite of this is e.g. a timer being used
	 * to guess when the display hardware has switched to the new image
	 * content.
	 */
	WP_PRESENTATION_FEEDBACK_KIND_HW_COMPLETION = 0x4,
	/**
	 * presentation was done zero-copy
	 *
	 * The presentation of this update was done zero-copy. This means
	 * the buffer from the client was given to display hardware as is,
	 * without copying it. Compositing with OpenGL counts as copying,
	 * even if textured directly from the client buffer. Possible
	 * zero-copy cases include direct scanout of a fullscreen surface
	 * and a surface on a hardware overlay.
	 */
	WP_PRESENTATION_FEEDBACK_KIND_ZERO_COPY = 0x8,
};
#endif /* WP_PRESENTATION_FEEDBACK_KIND_ENUM */

/**
 * @ingroup iface_wp_presentation_feedback
 * @struct wp_presentation_feedback_listener
 */
struct wp_presentation_feedback_listener {
	/**
	 * presentation synchronized to this output
	 *
	 * As presentation can be synchronized to only one output at a
	 * time, this event tells which output it was. This event is only
	 * sent prior to the presented event.
	 *
	 * As clients may bind to the same global wl_output multiple times,
	 * this event is sent for each bound instance that matches the
	 * synchronized output. If a client has not bound to the right
	 * wl_output global at all, this event is not sent.
	 * @param output presentation output
	 */
	void (*sync_output)(void *data,
			    struct wp_presentation_feedback *wp_presentation_feedback,
			    struct wl_output *output);
	/**
	 * the content update was displayed
	 *
	 * The associated content update was displayed to the user at the
	 * indicated time (tv_sec_hi/lo, tv_nsec). For the interpretation
	 * of the timestamp, see presentation.clock_id event.
	 *
	 * The timestamp corresponds to the time when the content update
	 * turned into light the first time on the surface's main output.
	 * Compositors may approximate this from the framebuffer flip
	 * completion events from the system, and the latency of the
	 * physical display path if known.
	 *
	 * This event is preceded by all related sync_output events telling
	 * which output's refresh cycle the feedback corresponds to, i.e.
	 * the main output for the surface. Compositors are recommended to
	 * choose the output containing the largest part of the wl_surface,
	 * or keeping the output they previously chose. Having a stable
	 * presentation output association helps clients predict future
	 * output refreshes (vblank).
	 *
	 * The 'refresh' argument gives the compositor's prediction of how
	 * many nanoseconds after tv_sec, tv_nsec the very next output
	 * refresh may occur. This is to further aid clients in predicting
	 * future refreshes, i.e., estimating the timestamps targeting the
	 * next few vblanks. If such prediction cannot usefully be done,
	 * the argument is zero.
	 *
	 * If the output does not have a constant refresh rate, explicit
	 * video mode switches excluded, then the refresh argument must be
	 * zero.
	 *
	 * The 64-bit value combined from seq_hi and seq_lo is the value of
	 * the output's vertical retrace counter when the content update
	 * was first scanned out to the display. This value must be
	 * compatible with the definition of MSC in GLX_OML_sync_control
	 * specification. Note, that if the display path has a non-zero
	 * latency, the time instant specified by this counter may differ
	 * from the timestamp's.
	 *
	 * If the output does not have a concept of vertical retrace or a
	 * refresh cycle, or the output device is self-refreshing without a
	 * way to query the refresh count, then the arguments seq_hi and
	 * seq_lo must be zero.
	 * @param tv_sec_hi high 32 bits of the seconds part of the presentation timestamp
	 * @param tv_sec_lo low 32 bits of the seconds part of the presentation timestamp
	 * @param tv_nsec nanoseconds part of the presentation timestamp
	 * @param refresh nanoseconds till next refresh
	 * @param seq_hi high 32 bits of refresh counter
	 * @param seq_lo low 32 bits of refresh counter
	 * @param flags combination of 'kind' values
	 */
	void (*presented)(void *data,
			  struct wp_presentation_feedback *wp_presentation_feedback,
			  uint32_t tv_sec_hi,
			  uint32_t tv_sec_lo,
			  uint32_t tv_nsec,
			  uint32_t refresh,
			  uint32_t seq_hi,
			  uint32_t seq_lo,
			  uint32_t flags);
	/**
	 * the content update was not displayed
	 *
	 * The content update was never displayed to the user.
	 */
	void (*discarded)(void *data,
			  struct wp_presentation_feedback *wp_presentation_feedback);
};

/**
 * @ingroup iface_wp_presentation_feedback
 */
static inline int
wp_presentation_feedback_add_listener(struct wp_presentation_feedback *wp_presentation_feedback,
				      const struct wp_presentation_feedback_listener *listener, void *data)
{
	return wl_proxy_add_listener((struct wl_proxy *) wp_presentation_feedback,
				     (void (**)(void)) listener, data);
}

/**
 * @ingroup iface_wp_presentation_feedback
 */
#define WP_PRESENTATION_FEEDBACK_SYNC_OUTPUT_SINCE_VERSION 1
/**
 * @ingroup iface_wp_presentation_feedback
 */
#define WP_PRESENTATION_FEEDBACK_PRESENTED_SINCE_VERSION 1
/**
 * @ingroup iface_wp_presentation_feedback
 */
#define WP_PRESENTATION_FEEDBACK_DISCARDED_SINCE_VERSION 1


/** @ingroup iface_wp_presentation_feedback */
static inline void
wp_presentation_feedback_set_user_data(struct wp_presentation_feedback *wp_presentation_feedback, void *user_data)
{
	wl_proxy_set_user_data((struct wl_proxy *) wp_presentation_feedback, user_data);
}

/** @ingroup iface_wp_presentation_feedback */
static inline void *
wp_presentation_feedback_get_user_data(struct wp_presentation_feedback *wp_presentation_feedback)
{
	return wl_proxy_get_user_data((struct wl_proxy *) wp_presentation_feedback);
}

static inline uint32_t
wp_presentation_feedback_get_version(struct wp_presentation_feedback *wp_presentation_feedback)
{
	return wl_proxy_get_version((struct wl_proxy *) wp_presentation_feedback);
}

/** @ingroup iface_wp_presentation_feedback */
static inline void
wp_presentation_feedback_destroy(struct wp_presentation_feedback *wp_presentation_feedback)
{
	wl_proxy_destroy((struct wl_proxy *) wp_presentation_feedback);
}

#ifdef  __cplusplus
}
#endif

#endif
