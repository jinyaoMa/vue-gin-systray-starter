module.exports = {
  preset: "@vue/cli-plugin-unit-jest/presets/typescript-and-babel",
  transform: {
    "^.+\\.vue$": "vue-jest",
  },
  testMatch: [
    "<rootDir>/client/tests/unit/**/*.spec.(js|jsx|ts|tsx)",
    "<rootDir>/client/**/__tests__/*.(js|jsx|ts|tsx)",
  ],
  moduleNameMapper: {
    "^@\\/(.*)$": "<rootDir>/client/$1",
  },
};
