// app/page.tsx - 在页面中使用 Header 组件
import Header from '@/components/Header';
import WeatherWidget from '@/components/WeatherWidget';

export default function HomePage() {
  return (
    <div>
      <Header />
      
      <main className="container mx-auto p-6">
        <h2 className="text-2xl font-bold mb-4">欢迎访问 Next.js 16</h2>
        <p className="text-gray-600">
          这个页面演示了客户端组件如何访问环境变量
        </p>
        <WeatherWidget />
      </main>
    </div>
  );
}