---
title: "PHP - Terbilang Rupiah - Number to Text"
description: "PHP - Terbilang Rupiah - Number to Text."
excerpt: "PHP - Terbilang Rupiah - Number to Text."
date: 2024-01-26T02:02:53+07:00
lastmod: 2024-01-26T02:02:57+07:00
draft: false
weight: 60
images: []
categories: ["Note"]
tags: ["php", "terbilang", "converter", "rupiah"]
contributors: ["Sugeng Sulistiyawan"]
pinned: false
homepage: false
---

## Terbilang Function PHP

```php
/**
 * Terbilang Rupiah
 *
 * @param float|int $nominal
 * @return string
 */
function terbilang($nominal)
{
    $f = new \NumberFormatter('id-ID', \NumberFormatter::SPELLOUT);

    return $f->format($nominal);
}
```

### EXAMPLE IMPLEMENTATION

```php
terbilang(1000) // seribu

terbilang(10000) // sepuluh ribu

terbilang(100) . ' rupiah' // seratus rupiah
```
