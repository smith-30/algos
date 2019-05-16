#include<cstdio>
#define N 301
using namespace std;
unsigned long long mod = 998244353LL, T1[N][N*N], T2[N][N*N/2];
     
int main()
{
  unsigned long long n,sum = 0,a[300], ans = 1, i,j;
  scanf("%llu",&n);
      
  for(i = 0; i < n ; ++i) {
    scanf("%llu",&a[i]);
    sum += a[i];
    ans = ans * 3 % mod;
  }
     
  T1[0][0] = 1;
  T2[0][0] = 1;
  for(i = 1 ; i <= n ; ++i) {
    for(j = 0; j <= sum ; ++j) {
      T1[i][j] = T1[i - 1][j] * 2 % mod;
      T2[i][j] = T2[i - 1][j];
          
      if(j >= a[i-1]) {
        T1[i][j] = (T1[i][j] + T1[i-1][j-a[i-1]]) % mod;
        T2[i][j] = (T2[i][j] + T2[i-1][j-a[i-1]]) % mod;
      }
    }
  }
     
  for(j = (sum+1)/2 ; j <= sum ; ++j)
    ans = (ans + mod - T1[n][j] * 3 % mod) % mod;
     
  if (sum%2 == 0)
    ans = (ans + T2[n][sum/2] * 3) % mod;
     
  printf("%llu",ans);
      
  return 0;
}