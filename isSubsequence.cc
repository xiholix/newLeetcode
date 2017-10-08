#include<iostream>

using namespace std;
class Solution {
public:
    bool isSubsequence(string s, string t) {
        const int s_len = 100;
        const int t_len = 100;
        bool sign[t_len][s_len] ;
        int s_len1=0; 
        s_len1 = s.length();
        int t_len1=0;
        t_len1  = t.length();
        // bool sign[t_len][s_len] = new bool[t_len][s_len];
        
        if (s.at(0)==t.at(0)){
            sign[0][0] = true;
        }
        
        for(int i=1; i<s_len; i++){
            sign[0][i] = false;
        }
        for(int i=1; i<t_len; i++){
            for(int j=0; j<s_len; j++)
                sign[i][j] = false;
        }

        for(int i=0; i<t_len; i++){
            for(int j=0; j<s_len; j++){
                if(t[i]==s[j] && sign[i-1][j-1]==true){
                    sign[i][j] = true;
                }else{
                    if(sign[i-1][j]==true)
                       sign[i][j]=true;
                }
            }
        }
        for(int i=0; i<6; i++){
            for(int j=0; j<3; j++){
               cout<<sign[i][j]<<"\t";
            }
            cout<<endl;
        }
        return sign[t_len1-1][s_len1-1];
    }
};

int main(){
    Solution t ;
    bool m = t.isSubsequence("abc", "ahbgdc");
    cout<<m<<endl;
    cout<<t.isSubsequence("axc", "ahbgdc")<<endl;
}