// app/api/data/route.ts - API Route 示例
import { NextResponse } from 'next/server';

export async function GET() {
  // 在服务器端，可以直接访问所有环境变量
  const dbUrl = process.env.DATABASE_URL;
  const openaiKey = process.env.OPENAI_API_KEY;
  
  // 这些私密信息永远不会暴露给浏览器
  console.log('数据库连接:', dbUrl); // ✅ 安全
  
  return NextResponse.json({ success: true });
}