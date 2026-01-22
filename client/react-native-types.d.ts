/// <reference types="react" />

declare module 'react-native' {
  import { ComponentType } from 'react';
  
  export const View: ComponentType<any>;
  export const Text: ComponentType<any>;
  export const StyleSheet: any;
  // Add other React Native components as needed
}
