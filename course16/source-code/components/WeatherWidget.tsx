// components/WeatherWidget.tsx - 客户端组件
'use client';

import { useState } from 'react';

interface WeatherData {
  city: string;
  temperature: number;
  condition: string;
  humidity: number;
  processedBy: string;
  timestamp: string;
  apiKeyUsed: string;
}

export default function WeatherWidget() {
  const [city, setCity] = useState('北京');
  const [weather, setWeather] = useState<WeatherData | null>(null);
  const [loading, setLoading] = useState(false);
  
  // ❌ 客户端无法访问私密环境变量
  const clientApiKey = process.env.WEATHER_API_KEY; // undefined
  console.log('客户端尝试访问 API Key:', clientApiKey); // undefined
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    
    try {
      // ✅ 客户端只调用内部 API，不直接接触 API Key
      const response = await fetch('/api/weather', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ city }),
      });
      
      const result = await response.json();
      
      if (result.success) {
        setWeather(result.data);
      } else {
        alert('查询失败');
      }
    } catch (error) {
      console.error('请求失败:', error);
      alert('网络错误');
    } finally {
      setLoading(false);
    }
  };
  
  return (
    <div className="max-w-md mx-auto p-6 bg-white rounded-lg shadow-lg">
      <h2 className="text-2xl font-bold mb-4">天气查询</h2>
      
      <form onSubmit={handleSubmit} className="mb-4">
        <input
          type="text"
          value={city}
          onChange={(e) => setCity(e.target.value)}
          placeholder="输入城市名称"
          className="w-full px-4 py-2 border rounded-lg mb-2"
        />
        <button
          type="submit"
          disabled={loading}
          className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 disabled:bg-gray-400"
        >
          {loading ? '查询中...' : '查询天气'}
        </button>
      </form>
      
      {weather && (
        <div className="bg-blue-50 p-4 rounded-lg">
          <h3 className="font-bold text-xl mb-2">{weather.city}</h3>
          <p className="text-3xl mb-1">{weather.temperature}°C</p>
          <p className="text-gray-600 mb-2">{weather.condition}</p>
          <p className="text-sm text-gray-500">湿度: {weather.humidity}%</p>
          <div className="mt-4 pt-4 border-t text-xs text-gray-400">
            <p>✅ 处理方式: {weather.processedBy}</p>
            <p>🔐 {weather.apiKeyUsed}</p>
            <p>🕐 {new Date(weather.timestamp).toLocaleString('zh-CN')}</p>
          </div>
        </div>
      )}
      
      <div className="mt-4 p-3 bg-yellow-50 border border-yellow-200 rounded text-sm">
        <p className="font-semibold mb-1">🔒 安全提示：</p>
        <p className="text-gray-700">
          客户端组件无法访问 <code className="bg-gray-200 px-1">WEATHER_API_KEY</code>
          （值为 undefined），但可以通过 API Route 安全地使用它。
          打开浏览器控制台查看日志验证！
        </p>
      </div>
    </div>
  );
}