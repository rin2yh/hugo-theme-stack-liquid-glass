+++
title = "SVG cover"
date = 2026-04-26
categories = ["Showcase"]
tags = ["svg"]
description = "Fixture confirming SVG cover artwork is served as-is without raster processing."
image = "cover.svg"
+++

This post pins down the SVG-cover code path: the cover image is passed through unchanged because Hugo's `.Resize` / `.Fill` are raster-only.
