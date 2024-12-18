package com.aiocw.novelmap

import android.content.Context
import android.content.DialogInterface
import android.os.Bundle
import android.util.Log
import android.view.View
import android.webkit.WebView
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat


class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main)
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main)) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }



        val myWebView: WebView = findViewById<WebView>(R.id.webview)
        myWebView.settings.javaScriptEnabled = true
        myWebView.settings.domStorageEnabled = true

        myWebView.webViewClient = MyWebViewClient()

//        myWebView.loadUrl("http://10.8.8.199:7780/ui/")
        myWebView.loadUrl("file:///android_asset/index.html");
    }

    fun writeUrl(url: String) {
        val sharedPref = getPreferences(Context.MODE_PRIVATE)
        with (sharedPref.edit()) {
            putString("url", url)
            apply()
        }
    }

    fun readUrl(): String {
        val sharedPref = getPreferences(Context.MODE_PRIVATE)
        val url = sharedPref.getString("url", "defaultName").toString()
        return url
    }
}