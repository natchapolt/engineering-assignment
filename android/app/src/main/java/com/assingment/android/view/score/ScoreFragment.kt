package com.assingment.android.view.score

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import com.assingment.android.R
import com.assingment.android.view.ShareViewModel

class ScoreFragment : Fragment() {

    private val layoutId: Int = R.layout.fragment_score

    private val model: ShareViewModel by viewModels()

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(layoutId, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val name = view.findViewById<TextView>(R.id.name)
        val score = view.findViewById<TextView>(R.id.score)
        name.text = model.name.value
        score.text = model.score.toString()
    }
}
