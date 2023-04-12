---
title: "Preloader HTML & CSS Style 1"
description: "Very Simple Loader use HTML & CSS only."
excerpt: "Very Simple Loader use HTML & CSS only."
date: 2023-04-12T06:19:59+07:00
lastmod: 2023-04-12T06:19:55+07:00
draft: false
weight: 60
images: ["thumbnail.png"]
categories: ["Preloader"]
tags: ["css", "html", "preloader", "animation"]
contributors: ["Sugeng Sulistiyawan"]
pinned: false
homepage: false
---

Preloader HTML & CSS Style 1.

Very Simple Loader use HTML & CSS only.

---

{{< youtube id="Hm6XQa0rz2w" autoplay="true" title="Preloader HTML & CSS Style 1" >}}

---

### CSS

```css
#preloader {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background-color: #d50000;
}

#preloader>.loader {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 146px;
  padding: 16px 40px;
  transform: translate(-50%, -50%);
}

#preloader>.loader>.dot,
#preloader>.loader>.dots>span {
  width: 24px;
  height: 24px;
  background-color: #fff;
}

#preloader>.loader>.dot {
  position: absolute;
  animation: dot 1.2s infinite;
  transform: translateX(0);
}

#preloader>.loader>.dots {
  animation: dots 1.2s infinite;
  transform: translateX(32px);
}

#preloader>.loader>.dots>span {
  display: block;
  float: left;
  margin-left: 8px;
  margin-right: 8px;
}

/* animation */
@keyframes dot {
  50% {
    transform: translateX(120px);
  }
}

@keyframes dots {
  50% {
    transform: translateX(-8px);
  }
}
```

### HTML

```html
<div id="preloader">
  <div class="loader">
    <span class="dot"></span>
    <div class="dots">
      <span></span>
      <span></span>
      <span></span>
    </div>
  </div>
</div>
```

#### Demo & Preview:

<script async src="//jsfiddle.net/sugengsulistiyawan/jhLftkuy/embed/result/"></script>
