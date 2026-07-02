/**
 * mobile-actions.js — speed-dial FAB (TOC + share) for small screens.
 *   Tapping the FAB fans out the actions. Share uses the Web Share API when
 *   available, falling back to an X/Twitter intent URL. TOC opens as a sheet.
 */
(function () {
  "use strict";

  function init() {
    var root = document.querySelector(".mobile-actions");
    if (!root) return;

    // --- Share (either the solo FAB or the speed-dial action) ---
    function doShare() {
      var title = root.getAttribute("data-share-title") || document.title;
      var url = root.getAttribute("data-share-url") || window.location.href;
      var fallback = root.getAttribute("data-share-fallback");
      if (navigator.share) {
        navigator.share({ title: title, url: url }).catch(function () {});
      } else if (fallback) {
        window.open(fallback, "_blank", "noopener,noreferrer");
      }
    }

    // --- Speed-dial menu ---
    var fabToggle = root.querySelector("[data-mobile-actions-toggle]");
    function closeMenu() {
      root.classList.remove("is-open");
      if (fabToggle) fabToggle.setAttribute("aria-expanded", "false");
    }
    if (fabToggle) {
      fabToggle.addEventListener("click", function (e) {
        e.stopPropagation();
        var open = root.classList.toggle("is-open");
        fabToggle.setAttribute("aria-expanded", open ? "true" : "false");
      });
      document.addEventListener("click", function (e) {
        if (root.classList.contains("is-open") && !root.contains(e.target)) {
          closeMenu();
        }
      });
    }

    Array.prototype.forEach.call(
      root.querySelectorAll("[data-mobile-share]"),
      function (btn) {
        btn.addEventListener("click", function () {
          closeMenu();
          doShare();
        });
      }
    );

    // --- TOC bottom sheet ---
    var sheet = document.getElementById("mobile-toc-sheet");
    var tocToggle = root.querySelector("[data-mobile-toc-toggle]");
    if (!sheet || !tocToggle) return;

    function openSheet() {
      closeMenu();
      sheet.classList.add("is-open");
      tocToggle.setAttribute("aria-expanded", "true");
      document.body.classList.add("sheet-open");
    }
    function closeSheet() {
      sheet.classList.remove("is-open");
      tocToggle.setAttribute("aria-expanded", "false");
      document.body.classList.remove("sheet-open");
    }

    tocToggle.addEventListener("click", openSheet);

    // One delegated handler: close on the overlay/close button or a TOC link.
    sheet.addEventListener("click", function (e) {
      if (
        e.target.closest("[data-mobile-toc-close]") ||
        e.target.closest('.mobile-toc__inner a[href^="#"]')
      ) {
        closeSheet();
      }
    });

    document.addEventListener("keydown", function (e) {
      if (e.key === "Escape") {
        closeMenu();
        closeSheet();
      }
    });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
