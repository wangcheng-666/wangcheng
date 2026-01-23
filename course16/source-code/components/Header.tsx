// components/Header.tsx - Client Component 示例
'use client';

export default function Header() {
  // ⚠️ 只有 NEXT_PUBLIC_ 开头的变量才能在客户端访问
  const siteName = process.env.NEXT_PUBLIC_SITE_NAME;
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  
  // ❌ 这会返回 undefined，因为没有 NEXT_PUBLIC_ 前缀
  const secret = process.env.NEXTAUTH_SECRET; // undefined
  
  // 在浏览器控制台打印，便于验证
  console.log('客户端可访问的变量:');
  console.log('✅ NEXT_PUBLIC_SITE_NAME:', siteName);
  console.log('✅ NEXT_PUBLIC_API_URL:', apiUrl);
  console.log('❌ NEXTAUTH_SECRET:', secret); // undefined
  
  return (
    <header className="bg-gradient-to-r from-blue-600 to-purple-600 text-white p-6 shadow-lg">
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mb-2">
          {siteName || '网站名称未设置'}
        </h1>
        <p className="text-blue-100 mb-4">
          API 地址: {apiUrl || 'API 地址未设置'}
        </p>
        
        <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4 text-sm">
          <p className="font-semibold mb-2">🔍 环境变量访问测试：</p>
          <div className="space-y-1 font-mono text-xs">
            <p>✅ NEXT_PUBLIC_SITE_NAME: {siteName ? '可访问' : '未定义'}</p>
            <p>✅ NEXT_PUBLIC_API_URL: {apiUrl ? '可访问' : '未定义'}</p>
            <p>❌ NEXTAUTH_SECRET: {secret ? '可访问' : 'undefined（正常）'}</p>
          </div>
          <p className="mt-3 text-xs text-blue-100">
            💡 打开浏览器控制台查看详细日志
          </p>
        </div>
      </div>
    </header>
  );
}