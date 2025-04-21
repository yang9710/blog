'use client';

import '@/styles/globals.css';
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuthStore } from '@/stores/auth';

export default function HomePage() {
  const router = useRouter();
  const { initialize } = useAuthStore();

  useEffect(() => {
    initialize();
    router.push('/articles');
  }, [initialize, router]);

  return null;
}