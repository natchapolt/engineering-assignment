package com.assingment.android.screen

import com.agoda.kakao.edit.KEditText
import com.agoda.kakao.screen.Screen
import com.agoda.kakao.text.KButton
import com.agoda.kakao.text.KTextView
import com.assingment.android.R

class HomeScreen : Screen<HomeScreen>() {
    val editText = KEditText { withId(R.id.nameInput) }
    val startButton = KButton { withId(R.id.nextButton) }
}
