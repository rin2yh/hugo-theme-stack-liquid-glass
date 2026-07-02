// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
  site: 'https://rin2yh.github.io',
  base: '/hugo-theme-stack-liquid-glass/',
  trailingSlash: 'always',
  integrations: [
    starlight({
      title: {
        en: 'Stack Liquid Glass — Docs',
        ja: 'Stack Liquid Glass — ドキュメント',
      },
      description: 'Usage documentation for the Stack Liquid Glass Hugo theme.',
      defaultLocale: 'root',
      locales: {
        root: { label: 'English', lang: 'en' },
        ja: { label: '日本語', lang: 'ja' },
      },
      sidebar: [
        {
          slug: 'getting-started',
          translations: { ja: 'はじめに' },
        },
        {
          slug: 'configuration',
          translations: { ja: '設定' },
        },
        {
          slug: 'og-images',
          translations: { ja: 'OGP 画像' },
        },
        {
          slug: 'widgets',
          translations: { ja: 'ウィジェット' },
        },
        {
          slug: 'shortcodes',
          translations: { ja: 'ショートコード' },
        },
        {
          slug: 'faq',
        },
      ],
      social: [
        {
          icon: 'github',
          label: 'GitHub',
          href: 'https://github.com/rin2yh/hugo-theme-stack-liquid-glass',
        },
      ],
    }),
  ],
});
