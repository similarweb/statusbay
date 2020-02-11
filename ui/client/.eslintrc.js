module.exports = {
    env: {
        browser: true,
        es6: true,
        jest: true
    },
    extends: [
        'plugin:react/recommended',
        'airbnb',
    ],
    globals: {
        Atomics: 'readonly',
        SharedArrayBuffer: 'readonly',
    },
    parserOptions: {
        ecmaFeatures: {
            jsx: true,
        },
        ecmaVersion: 2018,
        sourceType: 'module',
    },
    plugins: [
        'react',
    ],
    rules: {
        "radix": [0],
        "react/jsx-filename-extension": [1, {"extensions": [".js", ".jsx"]}],
        'import/no-extraneous-dependencies': [
            1,
            {
                devDependencies: [
                   '.storybook/**',
                    'src/**/*.stories.js',
                    'src/**/*.test.js',
                ]
            }
        ],
        'react/jsx-props-no-spreading': [0],
        'react/jsx-boolean-value': [0]
    },
    "overrides": [
        {
            "files": ["*.test.js"],
            "rules": {
                "no-unused-expressions": "off"
            }
        }
    ]
};
