"use client";

import { useEffect, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '@/store/store';
import { setUserData } from '@/store/userData';
import { authService } from '@/services/authService';
import { UserData } from '@/types/user';

export default function SessionInitializer() {
  const dispatch = useDispatch<AppDispatch>();
  const hasInitialized = useRef(false);

  useEffect(() => {
    // Only run once on app initialization
    if (hasInitialized.current) return;
    hasInitialized.current = true;

    const initializeSession = async () => {
      // Check if sessionId cookie exists
      const sessionId = document.cookie
        .split('; ')
        .find(row => row.startsWith('sessionId='))
        ?.split('=')[1];

      if (!sessionId) {
        return; // No session cookie, user needs to login
      }

      try {
        // Verify session and get user data
        const response = await authService.verifySession(sessionId);
        
        if (response.success) {
          // Store user_id in Redux store
          const userData: UserData = {
            user_id: (response.data as { user_id: number }).user_id as number
          };
          dispatch(setUserData(userData));
        }
      } catch (error) {
        console.log('Session verification failed:', error);
        // Session is invalid, user will need to login again
      }
    };

    initializeSession();
  }, [dispatch]);

  return null; // This component doesn't render anything
}
