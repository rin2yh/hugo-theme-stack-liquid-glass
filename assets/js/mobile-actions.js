/**
 * mobile-actions.js — floating TOC + share buttons for small screens.
 *   Share uses the Web Share API when available, falling back to an X/Twitter
 *   intent URL. The TOC opens as a bottom sheet.
 */
(function () {
  "use strict";

  function init() {
    var root = document.querySelector(".mobile-actions");
    if (!root) return;

    // --- Share ---
    var shareBtn = root.querySelector("[data-mobile-share]");
    if (shareBtn) {
      shareBtn.addEventListener("click", function () {
        var title = root.getAttribute("data-share-title") || document.title;
        var url = root.getAttribute("data-share-url") || window.location.href;
        var fallback = root.getAttribute("data-share-fallback");
        if (navigator.share) {
          navigator.share({ title: title, url: url }).catch(function () {});
        } else if (fallback) {
          window.open(fallback, "_blank", "noopener,noreferrer");
        }
      });
    }

    // --- TOC bottom sheet ---
    var sheet = document.getElementById("mobile-toc-sheet");
    var toggle = root.querySelector("[data-mobile-toc-toggle]");
    if (!sheet || !toggle) return;

    var closeTimer = null;

    function open() {
      if (closeTimer) {
        clearTimeout(closeTimer);
        closeTimer = null;
      }
      sheet.hidden = false;
      // Force reflow so the slide-up transition runs from the hidden state.
      void sheet.offsetWidth;
      sheet.classList.add("is-open");
      toggle.setAttribute("aria-expanded", "true");
      document.body.classList.add("sheet-open");
      document.addEventListener("keydown", onKey);
    }

    function close() {
      sheet.classList.remove("is-open");
      toggle.setAttribute("aria-expanded", "false");
      document.body.classList.remove("sheet-open");
      document.removeEventListener("keydown", onKey);
      closeTimer = setTimeout(function () {
        sheet.hidden = true;
      }, 400);
    }

    function onKey(e) {
      if (e.key === "Escape") close();
    }

    toggle.addEventListener("click", function () {
      if (sheet.hidden) open();
      else close();
    });

    Array.prototype.forEach.call(
      sheet.querySelectorAll("[data-mobile-toc-close]"),
      function (el) {
        el.addEventListener("click", close);
      }
    );

    // Dismiss the sheet once the reader jumps to a section.
    Array.prototype.forEach.call(
      sheet.querySelectorAll('.mobile-toc__inner a[href^="#"]'),
      function (a) {
        a.addEventListener("click", close);
      }
    );
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
