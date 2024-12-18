package com.aiocw.novelmap

import android.app.Dialog
import android.content.Context
import android.content.DialogInterface
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.LinearLayout
import android.widget.TextView

class CustomDialog : Dialog {

    constructor(context: Context?) : super(context!!)
    constructor(context: Context?, theme: Int) : super(context!!, theme)

    class Builder(private val context: Context) {
        var content: String? = null
        private var positiveButtonClickListener: DialogInterface.OnClickListener? = null

        fun setContent(content: String): Builder {
            this.content = content
            return this
        }

        fun setPositiveButton(listener: DialogInterface.OnClickListener): Builder {
            this.positiveButtonClickListener = listener
            return this
        }

        fun getEditeText() : String? {
            return  content
        }

        fun create(): CustomDialog {
            val layoutInflater = context.getSystemService(Context.LAYOUT_INFLATER_SERVICE) as LayoutInflater
            //为自定义弹框设置主题
            val customDialog = CustomDialog(context, R.style.Theme_NovelMap)
            val view = layoutInflater.inflate(R.layout.dialog_exit, null)
            customDialog.addContentView(view, LinearLayout.LayoutParams(
                LinearLayout.LayoutParams.FILL_PARENT,
                LinearLayout.LayoutParams.WRAP_CONTENT
            ))
            //设置弹框内容
            content?.let {
                (view.findViewById(R.id.dialog_content) as TextView).text = it
            }
            //设置弹框按钮
            positiveButtonClickListener?.let {
                (view.findViewById(R.id.dialog_sure) as Button).setOnClickListener {
                    content = (view.findViewById(R.id.editText) as EditText).text.toString()
                    Log.v("Daadd", content.toString())
                    positiveButtonClickListener!!.onClick(customDialog, DialogInterface.BUTTON_POSITIVE)

                }
            } ?: run {
                (view.findViewById(R.id.dialog_sure) as Button).visibility = View.GONE
            }
//            negativeButtonClickListener?.let {
//                (view.findViewById(R.id.editText) as EditText).setOnClickListener {
//                    negativeButtonClickListener!!.onClick(customDialog, DialogInterface.BUTTON_NEGATIVE)
//                }
//            } ?: run {
//                (view.findViewById(R.id.editText) as EditText).visibility = View.GONE
//            }
            customDialog.setContentView(view)
            return customDialog
        }
    }
}